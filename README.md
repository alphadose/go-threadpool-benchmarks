Benchmarks of itogami vs other go threadpools

Comparing against:-

1. Unlimited goroutines
2. [Ants](https://github.com/panjf2000/ants)
3. [Gamma-Zero-Worker-Pool](https://github.com/gammazero/workerpool)

Normal Pool
```bash
$ go test -bench=. -benchmem constants.go general_test.go
goos: darwin
goarch: arm64
BenchmarkUnlimitedGoroutines-8   	       4	 294670417 ns/op	96600320 B/op	 2004185 allocs/op
BenchmarkAntsPool-8              	       2	 510125334 ns/op	22748144 B/op	 1101798 allocs/op
BenchmarkGammaZeroPool-8         	       2	 697477271 ns/op	18863856 B/op	 1048377 allocs/op
BenchmarkItogamiPool-8           	       4	 317310260 ns/op	25431518 B/op	 1055384 allocs/op
PASS
ok  	command-line-arguments	10.764s
```

Pool With Func
```bash
go test -bench=. -benchmem constants.go throughput_test.go
goos: darwin
goarch: arm64
BenchmarkAntsPooWithFunc-8       	       2	 553801375 ns/op	 6429988 B/op	   98143 allocs/op
BenchmarkItogamiPoolWithFunc-8   	       3	 362104570 ns/op	15861378 B/op	   83403 allocs/op
PASS
ok  	command-line-arguments	4.334s
```
