package test

import (
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
