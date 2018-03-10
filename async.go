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
