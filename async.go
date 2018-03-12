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

func RunForever(concurrent int, task Task) {
	var wg sync.WaitGroup
	wg.Add(concurrent)

	for c := 0; c < concurrent; c++ {
		go func() {
			defer wg.Done()
			for {
				task()
			}
		}()
	}

	wg.Wait()
}

func RunLimited(concurrent int, count int, task Task) {
	var wg sync.WaitGroup
	wg.Add(concurrent)

	for c := 0; c < concurrent; c++ {
		go func() {
			defer wg.Done()
			for i := 0; i < count; i++ {
				task()
			}
		}()
	}

	wg.Wait()
}
