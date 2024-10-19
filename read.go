package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

// Name struct with fname and lname fields, each up to 20 characters
type Name struct {
    fname string
    lname string
}

func main() {
    // Prompt the user for the name of the text file
    var filename string
    fmt.Print("Enter the name of the text file: ")
    fmt.Scanln(&filename)

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer file.Close()

    // Slice to hold Name structs
    var names []Name

    // Read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Split the line into first name and last name
        fields := strings.Fields(line)
        if len(fields) != 2 {
            fmt.Printf("Ignoring line: %s\n", line)
            continue
        }

        // Truncate the names to 20 characters if necessary
        fname := truncateString(fields[0], 20)
        lname := truncateString(fields[1], 20)

        // Create a Name struct and add it to the slice
        name := Name{fname: fname, lname: lname}
        names = append(names, name)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    // Iterate through the slice and print the names
    for _, name := range names {
        fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
    }
}

// Helper function to truncate a string to a maximum length
func truncateString(s string, maxLength int) string {
    if len(s) > maxLength {
        return s[:maxLength]
    }
    return s
}
