package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a string")

	input,_ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	inputLower := strings.ToLower(input)

	// Check if the string starts with 'i', contains 'a', and ends with 'n'
    if strings.HasPrefix(inputLower, "i") &&
        strings.Contains(inputLower, "a") &&
        strings.HasSuffix(inputLower, "n") {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }

}