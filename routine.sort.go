package main

import (
	"fmt"
	"sort"
	"sync"
)

// Function to merge sorted subarrays
func merge(sortedArrays [][]int) []int {
	result := sortedArrays[0]
	for i := 1; i < len(sortedArrays); i++ {
		result = mergeTwo(result, sortedArrays[i])
	}
	return result
}

// Helper function to merge two sorted arrays
func mergeTwo(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}

func main() {
	// Prompt user for input
	var input int
	fmt.Println("Enter a series of integers (type -1 to end):")

	var array []int
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		array = append(array, input)
	}

	if len(array) == 0 {
		fmt.Println("No integers entered!")
		return
	}

	// Determine partition sizes
	partSize := (len(array) + 3) / 4 // Ensures that partition size is approximately equal
	sortedArrays := make([][]int, 4) // 4 subarrays for sorting

	var wg sync.WaitGroup
	wg.Add(4) // 4 goroutines

	// Function to sort a subarray in a goroutine
	sortSubarray := func(part []int, idx int) {
		defer wg.Done()
		fmt.Printf("Goroutine %d sorting subarray: %v\n", idx+1, part)
		sort.Ints(part)
		sortedArrays[idx] = part
		fmt.Printf("Goroutine %d sorted subarray: %v\n", idx+1, part)
	}

	// Launch 4 goroutines to sort different parts of the array
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(array) {
			end = len(array)
		}
		go sortSubarray(array[start:end], i)
	}

	wg.Wait() // Wait for all goroutines to complete

	// Merge sorted subarrays
	finalSortedArray := merge(sortedArrays)

	// Print final sorted array
	fmt.Println("Final sorted array:", finalSortedArray)
}
