// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ratelimit_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/joshdk/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ratelimit.NewPool(10, time.Second)
}

func TestStop(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Stop()
}

func TestWait(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Wait()
}

func TestMultiWait(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Wait()
	pool.Wait()
	pool.Wait()
}

func TestStopWait(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Stop()
	pool.Wait()
}

func TestMultiStop(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Stop()
	pool.Stop()
	pool.Stop()
}

func TestMultiStopWait(t *testing.T) {
	pool := ratelimit.NewPool(10, time.Second)

	pool.Stop()
	pool.Stop()
	pool.Stop()
	pool.Wait()
}

func TestWaitForAllTasksToExecute(t *testing.T) {

	pool := ratelimit.NewPool(1, time.Second)

	var ops uint64 = 0

	for i := 0; i < 10; i++ {
		pool.Add(func() {
			atomic.AddUint64(&ops, 1)
		})
	}

	// Wait for all executing tasks to finish
	pool.Wait()

	assert.True(t, 10 == ops)
}

func TestWaitForExecutingTasksToFinish(t *testing.T) {
	pool := ratelimit.NewPool(1, time.Second)

	var ops uint64 = 0

	for i := 0; i < 10; i++ {
		pool.Add(func() {
			time.Sleep(10 * time.Second)
			atomic.AddUint64(&ops, 1)
		})
	}

	// Allow time for the first 5 tasks to be executed
	time.Sleep(5 * time.Second)

	// Stop execution of any additional tasks
	pool.Stop()

	// Wait for all executing tasks to finish
	pool.Wait()

	assert.True(t, ops != 0)
	assert.True(t, ops != 10)
	assert.True(t, 5 <= ops)
	assert.True(t, ops <= 6)
}

func TestTasksShouldNotExecuteAfterWaiting(t *testing.T) {
	pool := ratelimit.NewPool(1, time.Second)

	pool.Wait()

	// Add 10 tasks that should never execute
	for i := 0; i < 10; i++ {
		pool.Add(func() {
			panic("This should never be called")
		})
	}

	pool.Wait()
}
