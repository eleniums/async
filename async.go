package async

import (
	"context"
	"sync"
)

// Task is a function that can be run concurrently.
type Task func() error

// Run will execute the given tasks concurrently and return any errors.
func Run(tasks ...Task) <-chan error {
	errc := make(chan error)

	// run tasks
	var wg sync.WaitGroup
	for _, v := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			err := task()
			if err != nil {
				errc <- err
			}
		}(v)
	}

	// make sure to close error channel
	go func() {
		wg.Wait()
		close(errc)
	}()

	return errc
}

// RunForever will execute the given task repeatedly on a set number of goroutines and return any errors. Context can be used to cancel execution of additional tasks.
func RunForever(ctx context.Context, concurrent int, task Task) <-chan error {
	errc := make(chan error)

	// run tasks
	var wg sync.WaitGroup
	for c := 0; c < concurrent; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				err := task()
				if err != nil {
					errc <- err
				}

				select {
				case <-ctx.Done():
					errc <- ctx.Err()
					return
				default:
				}
			}
		}()
	}

	// make sure to close error channel
	go func() {
		wg.Wait()
		close(errc)
	}()

	return errc
}

// RunLimited will execute the given task a set number of times on a set number of goroutines and return any errors. Total times the task will be executed is equal to concurrent multiplied by count. Context can be used to cancel execution of additional tasks.
func RunLimited(ctx context.Context, concurrent int, count int, task Task) <-chan error {
	errc := make(chan error)

	// run tasks
	var wg sync.WaitGroup
	for c := 0; c < concurrent; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < count; i++ {
				err := task()
				if err != nil {
					errc <- err
				}

				select {
				case <-ctx.Done():
					errc <- ctx.Err()
					return
				default:
				}
			}
		}()
	}

	// make sure to close error channel
	go func() {
		wg.Wait()
		close(errc)
	}()

	return errc
}

// Wait until channel is closed or error is received.
func Wait(errc <-chan error) error {
	for err := range errc {
		if err != nil {
			return err
		}
	}

	return nil
}

// HandleError sets a handler function to be called anytime an error is received on the given channel.
func HandleError(errc <-chan error, handler func(error)) {
	go func() {
		for err := range errc {
			if err != nil {
				handler(err)
			}
		}
	}()
}
