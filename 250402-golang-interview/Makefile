.PHONY: help bench bench-read bench-write bench-mixed bench-all test race clean fmt vet

help:
	@echo "Available targets:"
	@echo "  make bench         - Run all benchmarks"
	@echo "  make bench-read    - Run read-heavy benchmarks"
	@echo "  make bench-write   - Run write-heavy benchmarks"
	@echo "  make bench-mixed   - Run mixed workload benchmarks"
	@echo "  make bench-compare - Run benchmarks and save results for comparison"
	@echo "  make test          - Run all tests"
	@echo "  make race          - Run tests with race detector"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make clean         - Clean benchmark results"

# Run all benchmarks
bench:
	@echo "Running all benchmarks..."
	go test -bench=. -benchmem -benchtime=3s

# Run read-heavy benchmarks
bench-read:
	@echo "Running read-heavy benchmarks..."
	go test -bench=Read -benchmem -benchtime=3s

# Run write-heavy benchmarks
bench-write:
	@echo "Running write-heavy benchmarks..."
	go test -bench=Write -benchmem -benchtime=3s

# Run mixed workload benchmarks
bench-mixed:
	@echo "Running mixed workload benchmarks..."
	go test -bench=Mixed -benchmem -benchtime=3s

# Run benchmarks and save results for comparison
bench-compare:
	@echo "Running benchmarks and saving results..."
	go test -bench=. -benchmem -benchtime=3s | tee benchmark_results.txt
	@echo "Results saved to benchmark_results.txt"

# Run benchmarks with CPU profile
bench-cpu:
	@echo "Running benchmarks with CPU profiling..."
	go test -bench=. -benchmem -cpuprofile=cpu.prof -benchtime=3s
	@echo "CPU profile saved to cpu.prof"
	@echo "View with: go tool pprof cpu.prof"

# Run benchmarks with memory profile
bench-mem:
	@echo "Running benchmarks with memory profiling..."
	go test -bench=. -benchmem -memprofile=mem.prof -benchtime=3s
	@echo "Memory profile saved to mem.prof"
	@echo "View with: go tool pprof mem.prof"

# Run benchmarks with mutex profile
bench-mutex:
	@echo "Running benchmarks with mutex profiling..."
	go test -bench=. -benchmem -mutexprofile=mutex.prof -benchtime=3s
	@echo "Mutex profile saved to mutex.prof"
	@echo "View with: go tool pprof mutex.prof"

# Run all tests
test:
	@echo "Running tests..."
	go test -v

# Run tests with race detector
race:
	@echo "Running tests with race detector..."
	go test -v -race

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Clean benchmark results and profiles
clean:
	@echo "Cleaning benchmark results and profiles..."
	rm -f benchmark_results.txt cpu.prof mem.prof mutex.prof

# Run specific cache benchmark
bench-mutex:
	go test -bench=Mutex -benchmem -benchtime=3s

bench-rwmutex:
	go test -bench=RWMutex -benchmem -benchtime=3s

bench-sharded:
	go test -bench=Sharded -benchmem -benchtime=3s

bench-syncmap:
	go test -bench=SyncMap -benchmem -benchtime=3s

bench-spinlock:
	go test -bench=SpinLock -benchmem -benchtime=3s

bench-cow:
	go test -bench=COW -benchmem -benchtime=3s

bench-hybrid:
	go test -bench=Hybrid -benchmem -benchtime=3s

# Run all checks
check: fmt vet test
	@echo "All checks passed!"
