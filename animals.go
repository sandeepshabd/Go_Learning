package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Define the Animal struct with food, locomotion, and noise fields
type Animal struct {
    food       string
    locomotion string
    noise      string
}

// Eat method prints the food of the animal
func (a Animal) Eat() {
    fmt.Println(a.food)
}

// Move method prints the locomotion method of the animal
func (a Animal) Move() {
    fmt.Println(a.locomotion)
}

// Speak method prints the sound made by the animal
func (a Animal) Speak() {
    fmt.Println(a.noise)
}

func main() {
    // Initialize animals with their respective data
    animals := map[string]Animal{
        "cow":   {"grass", "walk", "moo"},
        "bird":  {"worms", "fly", "peep"},
        "snake": {"mice", "slither", "hsss"},
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        // Prompt for input
        fmt.Print("> ")

        // Read input from user
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        // Split input into two parts: animal and action
        parts := strings.Split(input, " ")
        if len(parts) != 2 {
            fmt.Println("Please provide valid input in the format: <animal> <action>")
            continue
        }

        animalName := parts[0]
        action := parts[1]

        // Get the animal from the map
        animal, exists := animals[animalName]
        if !exists {
            fmt.Println("Invalid animal. Please choose from cow, bird, or snake.")
            continue
        }

        // Perform the requested action
        switch action {
        case "eat":
            animal.Eat()
        case "move":
            animal.Move()
        case "speak":
            animal.Speak()
        default:
            fmt.Println("Invalid action. Please choose from eat, move, or speak.")
        }
    }
}
