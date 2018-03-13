package test

import (
	"errors"
	"sync"
	"testing"

	"github.com/eleniums/async"

	assert "github.com/stretchr/testify/require"
)

func Test_Run_Successful(t *testing.T) {
	// arrange
	count1 := 0
	task1 := func() error {
		count1++
		return nil
	}

	count2 := 0
	task2 := func() error {
		count2++
		return nil
	}

	count3 := 0
	task3 := func() error {
		count3++
		return nil
	}

	// act
	err := async.Run(task1, task2, task3)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 1, count1)
	assert.Equal(t, 1, count2)
	assert.Equal(t, 1, count3)
}

func Test_Run_Error(t *testing.T) {
	// arrange
	count1 := 0
	task1 := func() error {
		count1++
		return nil
	}

	count2 := 0
	task2 := func() error {
		count2++
		return nil
	}

	count3 := 0
	task3 := func() error {
		count3++
		return errors.New("task3")
	}

	// act
	err := async.Run(task1, task2, task3)

	// assert
	assert.Error(t, err)
	assert.Equal(t, 1, count3)
}

func Test_RunLimited_Successful(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		count++
		return nil
	}

	// act
	err := async.RunLimited(3, 4, task)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 12, count)
}

func Test_RunLimited_Error(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		if count < 8 {
			count++
			if count == 8 {
				return errors.New("error")
			}
		}

		return nil
	}

	// act
	err := async.RunLimited(3, 4, task)

	// assert
	assert.Error(t, err)
	assert.Equal(t, 8, count)
}

func Test_RunForever_Successful(t *testing.T) {
	// arrange
	var wg sync.WaitGroup
	wg.Add(12)
	count := 0
	task := func() error {
		if count < 12 {
			defer wg.Done()
			count++
		}

		return nil
	}

	// act
	go async.RunForever(1, task)
	wg.Wait()

	// assert
	assert.Equal(t, 12, count)
}

func Test_RunForever_Error(t *testing.T) {
	// arrange
	count := 0
	task := func() error {
		if count < 12 {
			count++
			if count == 12 {
				return errors.New("error")
			}
		}

		return nil
	}

	// act
	err := async.RunForever(1, task)

	// assert
	assert.Error(t, err)
	assert.Equal(t, 12, count)
}
