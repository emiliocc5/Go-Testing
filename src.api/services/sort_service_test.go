package services

import (
	"Go-Testing/src.api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)
//THIS IS AN INTEGRATION TEST

func TestSort(t *testing.T) {
	//Init
	elements := sort.GetElements(10)

	//Execution
	Sort(elements)

	//Validation
	assert.Equal(t, elements[0], 0)
	assert.Equal(t, elements[len(elements) - 1], 9)
}

func TestSortWithMoreThan10000(t *testing.T) {
	//Init
	elements := sort.GetElements(10001)

	//Execution
	Sort(elements)

	//Validation
	assert.Equal(t, elements[0], 0)
	assert.Equal(t, elements[len(elements) - 1], 10000)
}

func BenchmarkSort10K(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100K(b *testing.B) {
	elements := sort.GetElements(1000000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}


