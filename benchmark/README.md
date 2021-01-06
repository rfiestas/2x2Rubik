# Benchmarks

I added some go benchmark tests that i used to found the better performance, also some other stuff to monitor the memory and the cpu.

## Benchmark

```bash
go test -bench=. -benchmem                    # Run all benchmarks(also show memory allocations).
go test -bench=Solve -benchmem                # Run selected benchmark.
go test -bench=. -benchmem -test.benchtime 2s # Increment duration of benchmark.
go test -bench=. -benchmem -test.count 2      # run benchmarks n times.
```

Output

```bash
go test -bench=Solve/ -test.count 1 -test.benchtime 1s
goos: darwin
goarch: amd64
pkg: github.com/rfiestas/rubik2x2/benchmark
BenchmarkSolve/engine-8               933352          1753 ns/op
BenchmarkSolve/function-8             558852          2591 ns/op
PASS
ok      github.com/rfiestas/rubik2x2/benchmark  16.052s
```

## Go runtime.MemStats

From <https://golangcode.com/print-the-current-memory-usage/>

```go
// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
```

Output:

```bash
go run .
Alloc = 0 MiB   TotalAlloc = 0 MiB  Sys = 68 MiB    NumGC = 0
Solved in level:  6
Alloc = 5893 MiB    TotalAlloc = 7855 MiB   Sys = 6163 MiB  NumGC = 2
```

## Time

```bash
TIMEFMT='%J   %U  user %S system %P cpu %*E total'$'\n'\n'avg shared (code):         %X KB'$'\n'\n'avg unshared (data/stack): %D KB'$'\n'\n'total (sum):               %K KB'$'\n'\n'max memory:                %M '$MAX_MEMORY_UNITS''$'\n'\n'page faults from disk:     %F'$'\n'\n'other page faults:         %R'
```

Output:

```bash
time sleep 4
sleep 4   0.00s  user 0.00s system 0% cpu 4.010 total
avg shared (code):         0 KB
avg unshared (data/stack): 0 KB
total (sum):               0 KB
max memory:                556 KB
page faults from disk:     0
other page faults:         374
```

```bash
 /usr/bin/time -l go run .
 ````

Output:

```bash
/usr/bin/time -l go run .
Solved in level: 10
        0.89 real         1.12 user         0.38 sys
 142233600  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
     68692  page reclaims
       403  page faults
         0  swaps
         0  block input operations
         0  block output operations
         0  messages sent
         0  messages received
       855  signals received
       227  voluntary context switches
     14200  involuntary context switches
```
