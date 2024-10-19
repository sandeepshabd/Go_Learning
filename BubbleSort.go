package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Swap function swaps the elements at index i and i+1 in the slice
func Swap(slice []int, i int) {
    temp := slice[i]
    slice[i] = slice[i+1]
    slice[i+1] = temp
}

// BubbleSort function sorts the slice in place using the bubble sort algorithm
func BubbleSort(slice []int) {
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-1-i; j++ {
            if slice[j] > slice[j+1] {
                Swap(slice, j)
            }
        }
    }
}

func main() {
    fmt.Println("Enter up to 10 integers, separated by spaces:")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    fields := strings.Fields(input)

    var nums []int
    for i, field := range fields {
        if i >= 10 {
            break
        }
        num, err := strconv.Atoi(field)
        if err != nil {
            fmt.Printf("Invalid input '%s'. Please enter integers only.\n", field)
            return
        }
        nums = append(nums, num)
    }

    BubbleSort(nums)

    fmt.Println("Sorted integers:")
    for _, num := range nums {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
}
