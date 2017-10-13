[![License](https://img.shields.io/github/license/joshdk/ratelimit.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/ratelimit?status.svg)](https://godoc.org/github.com/joshdk/ratelimit)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/ratelimit)](https://goreportcard.com/report/github.com/joshdk/ratelimit)
[![CircleCI](https://circleci.com/gh/joshdk/ratelimit.svg?&style=shield)](https://circleci.com/gh/joshdk/ratelimit/tree/master)
[![CodeCov](https://codecov.io/gh/joshdk/ratelimit/branch/master/graph/badge.svg)](https://codecov.io/gh/joshdk/ratelimit)

# Ratelimit

⏳ Super simple rate-limited task pool

## Usage

```go
// Create a pool that can run 10 tasks per second
pool := ratelimit.NewPool(10, time.Second)

// Add 100 tasks to the pool
for i := 0; i < 100; i++ {
	i := i
	pool.Add(func() {
		fmt.Printf("this is job #%d\n", i)
	})
}

// Sleep for 10 seconds as tasks execute
time.Sleep(10 * time.Second)

// Stop task execution and wait
pool.Stop()
pool.Wait()
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.