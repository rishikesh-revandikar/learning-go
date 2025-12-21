package main

import (
	"fmt"
	"math"
)

func main(){
	// BMI Claculator
	var weight float64
	var height float64

	fmt.Println(("--------- BMI Calculator ---------"))
	
	fmt.Print("Enter your weight in kg (e.g. 75.5): ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Invalid input for weight.")
		return
	}

	fmt.Print("Enter your height in meters (e.g. 1.75): ")
	_, err= fmt.Scan(&height)
	if err != nil {
		fmt.Println("Invalid input for height.")
		return
	}

	bmi := weight / math.Pow(height,2)


	fmt.Printf("\n Your BMI is: %.2f\n",bmi)

	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Category: Normal weight")
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}