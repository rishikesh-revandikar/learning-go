package main

import (
	"math/rand"
	"fmt"
)


func main(){
	
	secretNumber := rand.Intn(100) + 1 // Generate number between 1 and 100
	var guess int
	attempts := 0

	fmt.Println("=== Guess the Number Game ===")
	fmt.Println("I have chosen a number between 1 and 100. Can you guess it?")

	// 2. The Loop: In Go, 'for' is used for all types of loops (while/do-while)
	for {
		attempts++
		fmt.Printf("Attempt %d: Enter your guess: ", attempts)
		
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		// 3. Hints & Correct Logic
		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			// 4. Winning condition
			fmt.Printf("🎉 Correct! You found it in %d attempts.\n", attempts)
			break // Exit the loop
		}
	}

}