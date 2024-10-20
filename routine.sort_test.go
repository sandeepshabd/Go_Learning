package main

import (
	"reflect"
	"testing"
)

// Test mergeTwo function with different scenarios
func TestMergeTwo(t *testing.T) {
	tests := []struct {
		arr1    []int
		arr2    []int
		expected []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{4, 5, 6}, []int{4, 5, 6}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 1, 2}, []int{1, 2, 2}, []int{1, 1, 1, 2, 2, 2}},
	}

	for _, test := range tests {
		result := mergeTwo(test.arr1, test.arr2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("mergeTwo(%v, %v) = %v; expected %v", test.arr1, test.arr2, result, test.expected)
		}
	}
}

// Test merge function with multiple sorted arrays
func TestMerge(t *testing.T) {
	tests := []struct {
		sortedArrays [][]int
		expected     []int
	}{
		{
			[][]int{
				{1, 4, 5},
				{2, 3, 6},
				{7, 8},
				{9, 10},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5},
				{6},
				{7, 8, 9, 10},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[][]int{
				{},
				{1},
				{2, 3, 4},
				{5, 6, 7},
			},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[][]int{
				{10, 20, 30},
				{15, 25, 35},
				{5, 50},
				{},
			},
			[]int{5, 10, 15, 20, 25, 30, 35, 50},
		},
	}

	for _, test := range tests {
		result := merge(test.sortedArrays)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("merge(%v) = %v; expected %v", test.sortedArrays, result, test.expected)
		}
	}
}
