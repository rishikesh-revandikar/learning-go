package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func processAny(values ...interface{}){
	for _, val := range values{
		fmt.Printf("%v : %T\n",val,val)
	}
}

func main(){
	// total := sum(1,2,3,4,5,6,7,8)

	nums := []int{1,2,3,4,5,6,7,8}
	total := sum(nums...)
	fmt.Println(total)

	// processAny(1,2,"risi",5.6)
}