package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Create a new reader for standard input
    reader := bufio.NewReader(os.Stdin)

    // Prompt the user to enter a name
    fmt.Print("Enter a name: ")
    nameInput, _ := reader.ReadString('\n')
    nameInput = strings.TrimSpace(nameInput)

    // Prompt the user to enter an address
    fmt.Print("Enter an address: ")
    addressInput, _ := reader.ReadString('\n')
    addressInput = strings.TrimSpace(addressInput)

    // Create a map and add the name and address
    person := map[string]string{
        "name":    nameInput,
        "address": addressInput,
    }

    // Marshal the map into a JSON object
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshalling to JSON:", err)
        return
    }

    // Print the JSON object
    fmt.Println(string(jsonData))
}
