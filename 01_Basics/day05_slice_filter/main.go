package main

import "fmt"

func main(){
	names := []string{"Alice", "Bob", "Charlie", "Demian", "Elizabeth", "Frank", "Gopher"}

	var longNames []string

	fmt.Println("Original List : ",names)

	fmt.Println("Filtering names longer than 5 characters")


	for _, name := range names {
		if len(name) > 5 {
			fmt.Printf("Found long name: %s (%d chars)\n",name,len(name))

			longNames = append(longNames, name)
		}
	}

	fmt.Println("---Filter complete----")
	fmt.Printf("Filtered List : %v\n",longNames)
}