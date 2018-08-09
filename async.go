package async

import (
	"sync"
)

// Task is a function that can be run concurrently.
type Task func() error

// Run will execute the given tasks concurrently and stop if a task returns an error.
func Run(tasks ...Task) error {
	errc := make(chan error)

	// run tasks
	var wg sync.WaitGroup
	for t := range tasks {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errc <- tasks[i]()
		}(t)
	}

	go func() {
		wg.Wait()
		close(errc)
	}()

	// check for errors
	for err := range errc {
		if err != nil {
			return err
		}
	}

	return nil
}

// RunForever will execute the given task repeatedly on a set number of goroutines and stop if a task returns an error.
func RunForever(concurrent int, task Task) error {
	errc := make(chan error)

	// run tasks
	for c := 0; c < concurrent; c++ {
		go func() {
			for {
				errc <- task()
			}
		}()
	}

	// check for errors
	for {
		err := <-errc
		if err != nil {
			return err
		}
	}
}

// RunLimited will execute the given task a set number of times on a set number of goroutines and stop if a task returns an error. Total times the task will be executed is equal to concurrent multiplied by count.
func RunLimited(concurrent int, count int, task Task) error {
	errc := make(chan error)

	// run tasks
	for c := 0; c < concurrent; c++ {
		go func() {
			for i := 0; i < count; i++ {
				errc <- task()
			}
		}()
	}

	// check for errors
	for i := 0; i < concurrent*count; i++ {
		err := <-errc
		if err != nil {
			return err
		}
	}

	return nil
}
