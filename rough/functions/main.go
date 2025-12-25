package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "js", "c"
}

func ProcessIt(fn func(a int) int) {
	fn(1)
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// result := add(3,5)
	// languages := getLanguages()

	// fn := func(a int) int {
	// 	return 2
	// }

	// ProcessIt(fn)
	fn := processIt()

	fmt.Println(fn(2))
}
