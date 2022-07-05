package test

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/alphadose/itogami"
	"github.com/bytedance/gopkg/util/gopool"
)

const benchmarkTimes = 10000

func DoCopyStack(a, b int) int {
	if b < 100 {
		return DoCopyStack(0, b+1)
	}
	return 0
}

func testFunc() {
	DoCopyStack(0, 0)
}

func testPanicFunc() {
	panic("test")
}

func TestPool(t *testing.T) {
	p := gopool.NewPool("test", 100, gopool.NewConfig())
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		p.Go(func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		})
	}
	wg.Wait()
	if n != 2000 {
		t.Error(n)
	}
}

func TestPoolPanic(t *testing.T) {
	p := gopool.NewPool("test", 100, gopool.NewConfig())
	p.Go(testPanicFunc)
}

func BenchmarkPool(b *testing.B) {
	config := gopool.NewConfig()
	config.ScaleThreshold = 1
	p := gopool.NewPool("benchmark", int32(runtime.GOMAXPROCS(0)), config)
	var wg sync.WaitGroup
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTimes)
		for j := 0; j < benchmarkTimes; j++ {
			p.Go(func() {
				testFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkGo(b *testing.B) {
	var wg sync.WaitGroup
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTimes)
		for j := 0; j < benchmarkTimes; j++ {
			go func() {
				testFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkItogami(b *testing.B) {
	var wg sync.WaitGroup
	p := itogami.NewPool(PoolSize)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTimes)
		for j := 0; j < benchmarkTimes; j++ {
			p.Submit(func() {
				testFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}
