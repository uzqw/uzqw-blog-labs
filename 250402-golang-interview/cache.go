package cache

import (
	"hash/fnv"
	"runtime"
	"sync"
	"sync/atomic"
)

type DiskStatus struct {
	ID     string
	Health int
	Temp   int
}

// 1. Basic Mutex Cache
type MutexCache struct {
	mu    sync.Mutex
	disks map[string]*DiskStatus
}

func NewMutexCache() *MutexCache {
	return &MutexCache{
		disks: make(map[string]*DiskStatus),
	}
}

func (c *MutexCache) Get(id string) *DiskStatus {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.disks[id]
}

func (c *MutexCache) Update(id string, status *DiskStatus) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.disks[id] = status
}

// 2. RWMutex Cache
type RWMutexCache struct {
	mu    sync.RWMutex
	disks map[string]*DiskStatus
}

func NewRWMutexCache() *RWMutexCache {
	return &RWMutexCache{
		disks: make(map[string]*DiskStatus),
	}
}

func (c *RWMutexCache) Get(id string) *DiskStatus {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.disks[id]
}

func (c *RWMutexCache) Update(id string, status *DiskStatus) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.disks[id] = status
}

// 3. Sharded Lock Cache
const ShardCount = 32

type ShardedCache struct {
	shards [ShardCount]struct {
		mu    sync.RWMutex
		disks map[string]*DiskStatus
	}
}

func NewShardedCache() *ShardedCache {
	c := &ShardedCache{}
	for i := 0; i < ShardCount; i++ {
		c.shards[i].disks = make(map[string]*DiskStatus)
	}
	return c
}

func (c *ShardedCache) getShard(id string) int {
	h := fnv.New32a()
	h.Write([]byte(id))
	return int(h.Sum32()) % ShardCount
}

func (c *ShardedCache) Get(id string) *DiskStatus {
	shard := &c.shards[c.getShard(id)]
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	return shard.disks[id]
}

func (c *ShardedCache) Update(id string, status *DiskStatus) {
	shard := &c.shards[c.getShard(id)]
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.disks[id] = status
}

// 4. sync.Map Cache
type SyncMapCache struct {
	disks sync.Map
}

func NewSyncMapCache() *SyncMapCache {
	return &SyncMapCache{}
}

func (c *SyncMapCache) Get(id string) *DiskStatus {
	v, ok := c.disks.Load(id)
	if !ok {
		return nil
	}
	return v.(*DiskStatus)
}

func (c *SyncMapCache) Update(id string, status *DiskStatus) {
	c.disks.Store(id, status)
}

// 5. Spinlock Cache
type SpinLockCache struct {
	lock  int32
	disks map[string]*DiskStatus
}

func NewSpinLockCache() *SpinLockCache {
	return &SpinLockCache{
		disks: make(map[string]*DiskStatus),
	}
}

func (c *SpinLockCache) Get(id string) *DiskStatus {
	// Spin acquire
	for !atomic.CompareAndSwapInt32(&c.lock, 0, 1) {
		runtime.Gosched() // Yield CPU to avoid starvation
	}
	// Very short critical section
	status := c.disks[id]
	atomic.StoreInt32(&c.lock, 0)
	return status
}

func (c *SpinLockCache) Update(id string, status *DiskStatus) {
	// Spin acquire
	for !atomic.CompareAndSwapInt32(&c.lock, 0, 1) {
		runtime.Gosched()
	}
	c.disks[id] = status
	atomic.StoreInt32(&c.lock, 0)
}

// 6. Copy-on-Write Cache
type COWCache struct {
	disks atomic.Value // stores map[string]*DiskStatus
}

func NewCOWCache() *COWCache {
	c := &COWCache{}
	c.disks.Store(make(map[string]*DiskStatus))
	return c
}

func (c *COWCache) Get(id string) *DiskStatus {
	m := c.disks.Load().(map[string]*DiskStatus)
	return m[id] // Read is completely lock-free!
}

func (c *COWCache) Update(id string, status *DiskStatus) {
	// Note: In production, you'd want a mutex here to prevent concurrent writers
	// from creating conflicting copies. For simplicity, we use Store directly.
	old := c.disks.Load().(map[string]*DiskStatus)
	// Copy entire map (write becomes slow)
	new := make(map[string]*DiskStatus, len(old)+1)
	for k, v := range old {
		new[k] = v
	}
	new[id] = status
	c.disks.Store(new)
}

// 7. Hybrid Cache (Sharded + COW)
type HybridCache struct {
	// Hot data: sharded lock protection
	hot [32]struct {
		mu   sync.RWMutex
		data map[string]*DiskStatus
	}
	// Cold data: COW (history records, rarely updated)
	cold atomic.Value
}

func NewHybridCache() *HybridCache {
	c := &HybridCache{}
	for i := 0; i < 32; i++ {
		c.hot[i].data = make(map[string]*DiskStatus)
	}
	c.cold.Store(make(map[string]*DiskStatus))
	return c
}

func (c *HybridCache) getShard(id string) int {
	h := fnv.New32a()
	h.Write([]byte(id))
	return int(h.Sum32()) % 32
}

func (c *HybridCache) Get(id string) *DiskStatus {
	// Try hot cache first
	shard := &c.hot[c.getShard(id)]
	shard.mu.RLock()
	status := shard.data[id]
	shard.mu.RUnlock()

	if status != nil {
		return status
	}

	// Fallback to cold cache
	m := c.cold.Load().(map[string]*DiskStatus)
	return m[id]
}

func (c *HybridCache) Update(id string, status *DiskStatus) {
	shard := &c.hot[c.getShard(id)]
	shard.mu.Lock()
	shard.data[id] = status
	shard.mu.Unlock()
}

func (c *HybridCache) UpdateCold(id string, status *DiskStatus) {
	// Note: In production, you'd want a mutex here to prevent concurrent writers
	old := c.cold.Load().(map[string]*DiskStatus)
	new := make(map[string]*DiskStatus, len(old)+1)
	for k, v := range old {
		new[k] = v
	}
	new[id] = status
	c.cold.Store(new)
}
