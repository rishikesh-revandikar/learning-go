package main

import "fmt"

func swap(a, b *int){
	temp := *a
	*a = *b
	*b = temp
}


func main()  {
	a, b := 1,2
	fmt.Printf("a : %v, b : %v\n",a,b)
	swap(&a, &b)
	fmt.Printf("a : %v, b : %v\n",a,b)
}