package main

import (
    "fmt"
)

func main() {
    var input float64

    fmt.Print("Enter a floating point number: ")
    _, err := fmt.Scan(&input)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid floating point number.")
        return
    }

    truncated := int(input)

    fmt.Println("Truncated integer:", truncated)
}
