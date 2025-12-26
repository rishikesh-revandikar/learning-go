package main

import "fmt"

func changeNum(num int){
	num = 5
	fmt.Println("In changeNum ", num)
}

func changeNumByRef(num *int){
	*num = 5 // derefrence

	fmt.Println("In changeNum ", *num)

}

func main(){
	num := 1
	
	// changeNum(num)
	changeNumByRef(&num)
	fmt.Println("Memory Address ", &num)
	fmt.Println("After changeNum in main",num)
}