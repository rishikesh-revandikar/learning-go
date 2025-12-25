package main

import "fmt"

func main(){
	// nums := []int {6,7,8}

	// for i := 0; i < len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	// for _, num := range nums{
	// 	fmt.Println(num)
	// }

	m := map[string]string {"fname":"Rishikesh", "lname":"Revandikar"}

	for k, v := range m{
		fmt.Println(k," : ",v)
	}

	// c contain unicode called Rune
	// i is starting byte of rune
	// 255 -> 1 byte
	// 300 -> 2 byte where i  points to first byte
	for i, c := range "golang"{
		fmt.Println(i,string(c))
	}
}