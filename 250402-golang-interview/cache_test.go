package cache

import (
	"fmt"
	"sync"
	"testing"
)

const (
	numKeys       = 1000
	readRatio     = 100 // 100:1 read:write ratio
	benchParallel = 32
)

// Prepare test data
func prepareTestData() []*DiskStatus {
	data := make([]*DiskStatus, numKeys)
	for i := 0; i < numKeys; i++ {
		data[i] = &DiskStatus{
			ID:     fmt.Sprintf("disk-%d", i),
			Health: 100,
			Temp:   45,
		}
	}
	return data
}

// Initialize cache with test data
func initMutexCache() *MutexCache {
	c := NewMutexCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initRWMutexCache() *RWMutexCache {
	c := NewRWMutexCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initShardedCache() *ShardedCache {
	c := NewShardedCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initSyncMapCache() *SyncMapCache {
	c := NewSyncMapCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initSpinLockCache() *SpinLockCache {
	c := NewSpinLockCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initCOWCache() *COWCache {
	c := NewCOWCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

func initHybridCache() *HybridCache {
	c := NewHybridCache()
	data := prepareTestData()
	for _, status := range data {
		c.Update(status.ID, status)
	}
	return c
}

// Benchmark: Read-heavy workload (100:1 read:write)
func BenchmarkMutexRead(b *testing.B) {
	c := initMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkRWMutexRead(b *testing.B) {
	c := initRWMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkShardedRead(b *testing.B) {
	c := initShardedCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkSyncMapRead(b *testing.B) {
	c := initSyncMapCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkSpinLockRead(b *testing.B) {
	c := initSpinLockCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkCOWRead(b *testing.B) {
	c := initCOWCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

func BenchmarkHybridRead(b *testing.B) {
	c := initHybridCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			c.Get(id)
			i++
		}
	})
}

// Benchmark: Write-heavy workload
func BenchmarkMutexWrite(b *testing.B) {
	c := initMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkRWMutexWrite(b *testing.B) {
	c := initRWMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkShardedWrite(b *testing.B) {
	c := initShardedCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkSyncMapWrite(b *testing.B) {
	c := initSyncMapCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkSpinLockWrite(b *testing.B) {
	c := initSpinLockCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkCOWWrite(b *testing.B) {
	c := initCOWCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

func BenchmarkHybridWrite(b *testing.B) {
	c := initHybridCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			status := &DiskStatus{ID: id, Health: 100, Temp: 45}
			c.Update(id, status)
			i++
		}
	})
}

// Benchmark: Mixed workload (100:1 read:write ratio)
func BenchmarkMutexMixed(b *testing.B) {
	c := initMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkRWMutexMixed(b *testing.B) {
	c := initRWMutexCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkShardedMixed(b *testing.B) {
	c := initShardedCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkSyncMapMixed(b *testing.B) {
	c := initSyncMapCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkSpinLockMixed(b *testing.B) {
	c := initSpinLockCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkCOWMixed(b *testing.B) {
	c := initCOWCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

func BenchmarkHybridMixed(b *testing.B) {
	c := initHybridCache()
	b.ResetTimer()
	b.SetParallelism(benchParallel)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			id := fmt.Sprintf("disk-%d", i%numKeys)
			if i%readRatio == 0 {
				status := &DiskStatus{ID: id, Health: 100, Temp: 45}
				c.Update(id, status)
			} else {
				c.Get(id)
			}
			i++
		}
	})
}

// Basic correctness tests
func TestCacheCorrectness(t *testing.T) {
	status := &DiskStatus{ID: "disk-1", Health: 100, Temp: 45}

	t.Run("MutexCache", func(t *testing.T) {
		c := NewMutexCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("RWMutexCache", func(t *testing.T) {
		c := NewRWMutexCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("ShardedCache", func(t *testing.T) {
		c := NewShardedCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("SyncMapCache", func(t *testing.T) {
		c := NewSyncMapCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("SpinLockCache", func(t *testing.T) {
		c := NewSpinLockCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("COWCache", func(t *testing.T) {
		c := NewCOWCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})

	t.Run("HybridCache", func(t *testing.T) {
		c := NewHybridCache()
		c.Update("disk-1", status)
		got := c.Get("disk-1")
		if got == nil || got.ID != "disk-1" {
			t.Errorf("expected disk-1, got %v", got)
		}
	})
}

// Concurrent correctness test
func TestCacheConcurrency(t *testing.T) {
	const goroutines = 100
	const operations = 1000

	t.Run("ShardedCache", func(t *testing.T) {
		c := NewShardedCache()
		var wg sync.WaitGroup
		wg.Add(goroutines)

		for i := 0; i < goroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < operations; j++ {
					key := fmt.Sprintf("disk-%d", j%100)
					status := &DiskStatus{ID: key, Health: 100, Temp: 45}
					c.Update(key, status)
					c.Get(key)
				}
			}(i)
		}

		wg.Wait()
	})
}
