package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		// 1. Display the Menu
		fmt.Println("\n--- CLI UTILITY MENU ---")
		fmt.Println("1. Greet Me")
		fmt.Println("2. Show Current Time")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error : Please enter valid choice.")

			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter your name: ")
			var name string
			fmt.Scan(&name)
			fmt.Printf("Hello, %s! Hope you're having a great day coding in Go.\n", name)
		case 2:
			// Go uses a specific reference date:
			// Mon Jan 2 15:04:05 MST 2006
			now := time.Now()
			fmt.Println("Current Time:", now.Format("15:04:05 PM"))
			fmt.Println("Full Date:", now.Format("Monday, 02-Jan-2006"))
		case 3:
			fmt.Println("Exiting... Goodbye!")
			os.Exit(0) // Successful exit
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}
