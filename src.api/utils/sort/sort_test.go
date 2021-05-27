package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//THIS IS A UNIT TEST

func TestBubbleSort(t *testing.T) {
	//Init
	elements := []int{9, 8, 6, 4, 27, 83}

	//Execution
	BubbleSort(elements)

	//Validation
	assert.Equal(t, elements[0], 4)
	assert.Equal(t, elements[len(elements) - 1], 83)
}


func TestSort(t *testing.T) {
	//Init
	elements := []int{9, 8, 6, 4, 27, 83}

	//Execution
	Sort(elements)

	//Validation
	assert.Equal(t, elements[0], 4)
	assert.Equal(t, elements[len(elements) - 1], 83)
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
