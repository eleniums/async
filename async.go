package async

import (
	"sync"
)

type Task func() error

func Run(tasks ...Task) error {
	count := len(tasks)

	errchan := make(chan error, count)

	for t := range tasks {
		go func(i int) {
			errchan <- tasks[i]()
		}(t)
	}

	for i := 0; i < count; i++ {
		err := <-errchan
		if err != nil {
			return err
		}
	}

	return nil
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
