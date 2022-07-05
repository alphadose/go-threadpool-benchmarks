# Benchmarks of itogami vs other go threadpools

Comparing against:-

1. Unlimited goroutines
2. [Ants](https://github.com/panjf2000/ants)
3. [Gamma-Zero-Worker-Pool](https://github.com/gammazero/workerpool)
4. [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
5. [Bytedance GoPool](https://github.com/bytedance/gopkg/tree/develop/util/gopool)

## Normal Pool

Results computed from [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) of 30 cases each via `go test -bench=. -benchmem constants.go general_test.go`

```bash
name                   time/op
UnlimitedGoroutines-8   301ms ± 4%
ErrGroup-8              515ms ± 9%
AntsPool-8              582ms ± 9%
GammaZeroPool-8         740ms ±13%
BytedanceGoPool-8       572ms ±18%
ItogamiPool-8           331ms ± 7%

name                   alloc/op
UnlimitedGoroutines-8  96.3MB ± 0%
ErrGroup-8              120MB ± 0%
AntsPool-8             22.4MB ± 6%
GammaZeroPool-8        18.8MB ± 1%
BytedanceGoPool-8      82.2MB ± 2%
ItogamiPool-8          25.6MB ± 2%

name                   allocs/op
UnlimitedGoroutines-8   2.00M ± 0%
ErrGroup-8              3.00M ± 0%
AntsPool-8              1.10M ± 2%
GammaZeroPool-8         1.05M ± 0%
BytedanceGoPool-8       2.59M ± 1%
ItogamiPool-8           1.05M ± 0%
```

## Pool With Predefined Function

```bash
go test -bench=. -benchmem constants.go throughput_test.go
goos: darwin
goarch: arm64
BenchmarkAntsPooWithFunc-8       	       2	 579326562 ns/op	 6581620 B/op	  102456 allocs/op

BenchmarkItogamiPoolWithFunc-8   	       3	 361101097 ns/op	14577853 B/op	   83407 allocs/op

PASS
ok  	command-line-arguments	4.334s
```
