package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sentence to analyze:")

	input , _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(strings.ToLower(input))

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	fmt.Println("\n--- Word Frequency ---")
	for word, count := range counts {
		fmt.Printf("%-12s : %d\n", word, count)
	}
}