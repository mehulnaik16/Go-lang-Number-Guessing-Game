package main

import (
	"fmt"
	"math/rand"
)

func game(chances int) {
	randomNumber := rand.Intn(100) + 1
	var choice int
	for i := 0; i < chances; i++ {
		fmt.Println("Enter your choice:")
		fmt.Scan(&choice)
		if randomNumber == choice {
			fmt.Println("You won the game!!!")
			return
		} else if randomNumber > choice {
			fmt.Printf("Incorrect! The number is greater than %d.\n", choice)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", choice)
		}
	}
	fmt.Println("You Lost The game!!!")
}

func main() {
	var choice int
	var playAgain string
	difficulties := []struct {
		name    string
		chances int
	}{
		{"Easy", 10},
		{"Medium", 5},
		{"Hard", 3},
	}

	for {
		fmt.Println("Select difficulty: 1. Easy 2. Medium 3. Hard")
		fmt.Scan(&choice)

		if choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		fmt.Printf("You selected %s difficulty. Let's start!\n", difficulties[choice-1].name)
		game(difficulties[choice-1].chances)

		fmt.Println("Do you want to play again? (y/n)")
		fmt.Scan(&playAgain)
		if playAgain != "y" && playAgain != "Y" {
			break
		}
	}

	fmt.Println("Thanks for playing!")
}
