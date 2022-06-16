package test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Jorropo/generator"
	"github.com/alphadose/itogami"
	"github.com/gammazero/workerpool"
	"github.com/panjf2000/ants/v2"
)

func demoFunc() {
	time.Sleep(time.Duration(BenchParam) * time.Millisecond)
}

func BenchmarkUnlimitedGoroutines(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			go func() {
				demoFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkAntsPool(b *testing.B) {
	var wg sync.WaitGroup
	p, _ := ants.NewPool(PoolSize, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.Release()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkGammaZeroPool(b *testing.B) {
	var wg sync.WaitGroup
	p := workerpool.New(PoolSize)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkItogamiPool(b *testing.B) {
	var wg sync.WaitGroup
	p := itogami.NewPool(PoolSize)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkGenerator(b *testing.B) {
	for i := b.N; i != 0; i-- {
		var count uint64 = 1<<64 - 1 // start at -1 since atomic.AddUint64 returns the new value (not the old one)
		p := generator.NewPool(func() (generator.Runner, bool) {
			i := atomic.AddUint64(&count, 1)
			if i < RunTimes {
				return demoFunc, true
			}
			return nil, false
		}, PoolSize)

		p.Wait()
	}
}
