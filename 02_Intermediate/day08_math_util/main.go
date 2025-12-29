package main

import "fmt"

func add(a,b int) (sum int){
	sum = a + b
	return
}

func subtract(a,b int) (diff int){
	diff = a - b
	return
}

func avg(nums ...int) (a float64){
	count := 0
	acc := 0
	for _, num := range nums {
		acc += num
		count++
	}

	a = float64(acc/count)
	return
}

func main()  {
	s := add(2,3)
	d := subtract(3,2)
	a := avg([]int{1,2,3,4,5}...)

	fmt.Printf("Sum: %d\nDifference: %d\nAverage: %.2f\n", s, d, a)
}