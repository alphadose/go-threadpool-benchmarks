package test

import (
	"sync"
	"testing"
	"time"

	"github.com/alphadose/itogami"
	"github.com/panjf2000/ants/v2"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/gammazero/workerpool"
	"golang.org/x/sync/errgroup"
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

func BenchmarkErrGroup(b *testing.B) {
	var wg sync.WaitGroup
	var pool errgroup.Group
	pool.SetLimit(PoolSize)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			pool.Go(func() error {
				demoFunc()
				wg.Done()
				return nil
			})
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

func BenchmarkBytedanceGoPool(b *testing.B) {
	var wg sync.WaitGroup
	p := gopool.NewPool("test", PoolSize, gopool.NewConfig())

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			p.Go(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
	b.StopTimer()
}
