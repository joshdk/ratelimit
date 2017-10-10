package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Task func()

type Pool struct {
	limiter *rate.Limiter
	ctx     context.Context
	cancel  context.CancelFunc
	tasks   chan Task
	group   sync.WaitGroup
}

func (pool *Pool) Add(task Task) {
	go func() {
		pool.tasks <- task
	}()
}

func (pool *Pool) Start() {
	pool.ctx, pool.cancel = context.WithCancel(context.Background())

	pool.group.Add(1)
	go func() {
		defer pool.group.Done()

		for {
			if err := pool.limiter.Wait(pool.ctx); err != nil {
				return
			}

			select {
			case <-pool.ctx.Done():
				return

			case task := <-pool.tasks:
				pool.group.Add(1)
				go func() {
					defer pool.group.Done()

					task()
				}()

			}
		}

	}()
}

func (pool *Pool) Stop() {
	pool.cancel()
}

func (pool *Pool) Wait() {
	pool.group.Wait()
}

func NewPool(actions float64, duration time.Duration) Pool {

	var limit = rate.Limit(actions / duration.Seconds())

	pool := Pool{
		limiter: rate.NewLimiter(limit, 1),
		tasks:   make(chan Task),
	}

	return pool
}
