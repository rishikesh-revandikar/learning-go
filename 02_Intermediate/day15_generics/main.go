package main

import "fmt"

func Reverse[T interface{}](s []T){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// Swap the elements
		s[i], s[j] = s[j], s[i]
	}
}

func main()  {
	// Testing with integers
	intSlice := []int{1, 2, 3, 4, 5}
	Reverse(intSlice)
	fmt.Println("Reversed Ints:   ", intSlice)

	// Testing with strings
	stringSlice := []string{"Go", "Generics", "Are", "Powerful"}
	Reverse(stringSlice)
	fmt.Println("Reversed Strings:", stringSlice)
}