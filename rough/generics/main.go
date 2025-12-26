package main

import "fmt"

// func printSlice[T interface{}](items []T)  {
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | string](items []T)  {
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T comparable](items []T)  {
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable, V string](items []T, name V)  {
	for _, item := range items{
		fmt.Println(item, name)
	}
}

type Stack[T interface{}] struct{
	elements []T
}

func main()  {
	// printSlice([]int{1,2,3,4})

	names := []string{"q","ds","we"}

	// printSlice(names)

	// stack := Stack{
	// 	elements: []int{1,2,3},
	// }

	// stack := Stack[string]{
	// 	elements: []string{"f","ds","ew"},
	// }

	printSlice(names,"rishi")
}