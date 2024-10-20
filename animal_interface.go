package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface that defines the Eat, Move, and Speak methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type implements the Animal interface
type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird type implements the Animal interface
type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake type implements the Animal interface
type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// Map to store the created animals
	animals := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		// Print prompt and read user input
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Split input into parts
		parts := strings.Split(input, " ")

		// Ensure we have exactly three parts
		if len(parts) != 3 {
			fmt.Println("Invalid input. Please provide the correct command format.")
			continue
		}

		command := parts[0]
		name := parts[1]
		info := parts[2]

		switch command {
		case "newanimal":
			// Create a new animal based on the type provided
			switch info {
			case "cow":
				animals[name] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[name] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[name] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal type. Choose from cow, bird, or snake.")
			}
		case "query":
			// Query information about an existing animal
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query. Choose from eat, move, or speak.")
			}
		default:
			fmt.Println("Invalid command. Choose from newanimal or query.")
		}
	}
}
