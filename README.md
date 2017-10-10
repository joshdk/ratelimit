[![License](https://img.shields.io/github/license/joshdk/ratelimit.svg)](https://opensource.org/licenses/MIT)

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

// Start scheduling tasks
pool.Start()

// Sleep for 10 seconds as tasks execute
time.Sleep(10 * time.Second)

// Stop task execution and wait
pool.Stop()
pool.Wait()
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.