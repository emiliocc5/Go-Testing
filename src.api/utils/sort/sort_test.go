package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//THIS IS A UNIT TEST

func TestBubbleSortOrderDesc(t *testing.T) {
	//Init
	elements := []int{9, 8, 6, 4, 27, 83}

	//Execution
	BubbleSort(elements)

	//Validation
	assert.Equal(t, elements[0], 83)
	assert.Equal(t, elements[len(elements) - 1], 4)
}
