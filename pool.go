// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

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
	group   sync.WaitGroup
}

// Add submits task into the pool for immediate (but rate-limited) execution.
// In the event that Stop() is called, this task will be discarded.
func (pool *Pool) Add(task Task) {
	pool.group.Add(1)
	go func() {
		defer pool.group.Done()

		if err := pool.limiter.Wait(pool.ctx); err != nil {
			return
		}

		task()
	}()
}

// Stop cancels all tasks in the pool, and prevents any others from being added.
func (pool *Pool) Stop() {
	pool.cancel()
}

// Wait blocks until either: All tasks in the pool have finished executing, or Stop() has been called.
// Tasks that are currently executing are allowed to finish
func (pool *Pool) Wait() {
	go func() {
		pool.group.Wait()
		pool.Stop()
	}()

	<-pool.ctx.Done()
	pool.group.Wait()
}

func NewPool(actions float64, duration time.Duration) *Pool {

	var limit = rate.Limit(actions / duration.Seconds())

	pool := Pool{
		limiter: rate.NewLimiter(limit, 1),
	}

	pool.ctx, pool.cancel = context.WithCancel(context.Background())

	return &pool
}
