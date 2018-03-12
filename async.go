package async

import (
	"sync"
)

type Task func()

func Run(tasks ...Task) {
	count := len(tasks)

	var wg sync.WaitGroup
	wg.Add(count)

	for t := range tasks {
		go func(i int) {
			defer wg.Done()
			tasks[i]()
		}(t)
	}

	wg.Wait()
}

func RunForever(tasks ...Task) {
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)

	count := len(tasks)

	var wg sync.WaitGroup
	wg.Add(count)

	for t := range tasks {
		go func(i int) {
			defer wg.Done()
			for {
				tasks[i]()
			}
		}(t)
	}

	wg.Wait()
}

func RunLimited(count int, tasks ...Task) {
	var wg sync.WaitGroup
	wg.Add(count)

	for t := range tasks {
		go func(i int) {
			for {
				tasks[i]()
				wg.Done()
			}
		}(t)
	}

	wg.Wait()
}
