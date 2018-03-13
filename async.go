package async

import (
	"sync"
)

// Task is a function that can be run concurrently.
type Task func() error

// Run will execute the given tasks concurrently and stop if a task returns an error.
func Run(tasks ...Task) error {
	concurrent := len(tasks)
	errchan := make(chan error, concurrent)

	// run tasks
	for t := range tasks {
		go func(i int) {
			errchan <- tasks[i]()
		}(t)
	}

	// check for errors
	for i := 0; i < concurrent; i++ {
		err := <-errchan
		if err != nil {
			return err
		}
	}

	return nil
}

// RunForever will execute the given task repeatedly on a set number of goroutines and stop if a task returns an error.
func RunForever(concurrent int, task Task) error {
	errchan := make(chan error, concurrent)

	// run tasks
	for c := 0; c < concurrent; c++ {
		go func() {
			for {
				errchan <- task()
			}
		}()
	}

	// check for errors
	for {
		err := <-errchan
		if err != nil {
			return err
		}
	}
}

// RunLimited will execute the given task a set number of times on a set number of goroutines and stop if a task returns an error. Total times the task will be executed is equal to concurrent multiplied by count.
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
