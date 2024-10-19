package main

import(
	"fmt"
	"sort"
	"strconv"
)

func main(){
	var input string 
	slice := make([]int ,0,3)

	for{
	fmt.Print("Enter an integer (or 'X' to exit): ")
    fmt.Scan(&input)
	if input == "X" || input == "x" {
            break
        }

        // Convert the input string to an integer
        num, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter an integer or 'X' to exit.")
            continue
        }

        // Append the integer to the slice
        slice = append(slice, num)

        // Sort the slice
        sort.Ints(slice)

        // Print the sorted slice
        fmt.Println("Sorted slice:", slice)
	}
}