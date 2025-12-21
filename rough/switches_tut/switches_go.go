package main

import "fmt"

func main(){
	
	TypeCheck := func(n interface{}){
		switch i := n.(type){
		case int:
			fmt.Printf("\n %v is Integer type\n",i)
		case string:
			fmt.Printf("\n %v is String type\n",i)
		case bool:
			fmt.Printf("\n %v is Boolean type\n",i)
		default:
			fmt.Printf("\n %v is Other type\n",i)
		}
	}

	TypeCheck(45.09)
}