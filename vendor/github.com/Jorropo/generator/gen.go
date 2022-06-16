package generator

import "sync"

type Runner func()

type Generator func() (r Runner, ok bool)

type Pool struct {
	wg sync.WaitGroup
	g  Generator
}

func NewPool(g Generator, count int) *Pool {
	p := &Pool{g: g}

	p.wg.Add(count)
	for ; count != 0; count-- {
		go p.pump()
	}

	return p
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) pump() {
	defer p.wg.Done()

	g := p.g

	for {
		task, ok := g()
		if !ok {
			break
		}
		task()
	}
}
