package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(10) + 1
	fmt.Println("Welcome to the Number Guesser!")
	fmt.Println("I'm thinking of a number between 1 and 10.")
	fmt.Println("Can you guess it?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess < 1 || guess > 10 {
			fmt.Println("Invalid guess. Please enter a number between 1 and 10.")
			continue
		}

		if guess == target {
			fmt.Println("Congratulations! You guessed the number.")
			break
		} else if guess < target {
			fmt.Println("Too low. Try again.")
		} else {
			fmt.Println("Too high. Try again.")
		}
	}
}
