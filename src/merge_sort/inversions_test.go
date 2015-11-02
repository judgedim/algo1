package merge_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"reader"
)

func TestCountInversionsReversed(t *testing.T) {
	array := []int64{6, 5, 4, 3, 2, 1}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 15, inversions)
}

func TestCountInversionsEqual(t *testing.T) {
	array := []int64{3, 3, 3, 3, 3, 3}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 0, inversions)
}

func TestCountInversionsSplit(t *testing.T) {
	array := []int64{1, 3, 5, 2, 4, 6}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 3, inversions)
}

func TestCountInversionsMixed(t *testing.T) {
	array := []int64{1, 7, 6, 3, 3, 5}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 7, inversions)
}

func TestCountInversionsSorted(t *testing.T) {
	array := []int64{1, 2, 3, 4, 5, 6}
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 0, inversions)
}

func TestCountInversions(t *testing.T) {
	array := reader.ReadArray("../../data/IntegerArray.txt")
	inversions, _ := countInversions(0, array)
	assert.Equal(t, 2407905288, inversions)
}

