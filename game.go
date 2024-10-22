package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

// Function that checks if guess is correct
func guess(guess, ans int) bool {
	switch {
	case guess < ans:
		fmt.Println("Incorrect! The number is greater than", guess)
		return false
	case guess > ans:
		fmt.Println("Incorrect! The number is less than", guess)
		return false
	default:
		return true
	}
}

// Function that plays one game round
func gameRound() {

	num := rand.IntN(100) + 1
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	// Difficulty level
	var dif int
	fmt.Print("Your choice>> ")
	_, err := fmt.Scan(&dif)
	if err != nil {
		log.Fatal("Incorrect usage!!!")
	}

	// Variable that indicates victory(true) or loss(false)
	wflag := false
	var choice int
	var numOfTries int
	switch dif {
	case 1:
		fmt.Println("Great! You have selected the Easy difficulty level.")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		for i := 1; i <= 10; i++ {
			fmt.Print("Your guess>> ")
			_, err := fmt.Scan(&choice)

			if err != nil {
				log.Fatal(err)
			}
			if guess(choice, num) {
				wflag = true
				numOfTries = i
				break
			}
		}
	case 2:
		fmt.Println("Great! You have selected the Medium difficulty level.")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		for i := 1; i <= 5; i++ {
			fmt.Print("Your guess>> ")
			_, err := fmt.Scan(&choice)

			if err != nil {
				log.Fatal(err)
			}
			if guess(choice, num) {
				wflag = true
				numOfTries = i
				break
			}
		}
	case 3:
		fmt.Println("Great! You have selected the Hard difficulty level.")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		for i := 1; i <= 3; i++ {
			fmt.Print("Your guess>> ")
			_, err := fmt.Scan(&choice)
			if err != nil {
				log.Fatal(err)
			}
			if guess(choice, num) {
				wflag = true
				numOfTries = i
				break
			}
		}
	default:
		log.Fatal("Incorrect usage!!!")

	}
	if wflag {
		fmt.Printf("Congratulations! You guessed the correct number in %d attempts", numOfTries)
	} else {
		fmt.Printf("You didn't guess :<")
	}
}

func main() {
	fmt.Println("------------------------------------")
	fmt.Println("Welcome to the Number Guessing Game!")
	gameRound()
}
