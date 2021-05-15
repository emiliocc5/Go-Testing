package services

import (
	"Go-Testing/src.api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)
//THIS IS AN INTEGRATION TEST

func TestSort(t *testing.T) {
	//Init
	elements := []int{9, 8, 6, 4, 27, 83}

	//Execution
	Sort(elements)

	//Validation
	assert.Equal(t, elements[0], 4)
	assert.Equal(t, elements[len(elements) - 1], 83)
}

func TestBubbleSort(t *testing.T) {
	//Init
	elements := []int{9, 8, 6, 4, 27, 83}

	//Execution
	BubbleSort(elements)

	//Validation
	assert.Equal(t, elements[0], 4)
	assert.Equal(t, elements[len(elements) - 1], 83)
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 8, 6, 4, 27, 83}
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 8, 6, 4, 27, 83}
	for i := 0; i < b.N; i++ {
		sort.Sort(elements)
	}
}
