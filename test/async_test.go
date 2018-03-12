package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/eleniums/async"

	assert "github.com/stretchr/testify/require"
)

func Test_Run_Successful(t *testing.T) {
	// arrange
	count1 := 0
	task1 := func() {
		count1++
	}

	count2 := 0
	task2 := func() {
		count2++
	}

	count3 := 0
	task3 := func() {
		count3++
	}

	// act
	async.Run(task1, task2, task3)

	// assert
	assert.Equal(t, 1, count1)
	assert.Equal(t, 1, count2)
	assert.Equal(t, 1, count3)
}

func Test_Async(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	isDone := false

	n := 0
	go func() {
		for {
			if isDone {
				break
			}

			n++
			fmt.Printf("%v\n", n)

			if n == 10 {
				fmt.Println("cancelling...")
				cancel()
			}
		}
	}()

	done := ctx.Done()
	<-done
	fmt.Println("after done")
}
