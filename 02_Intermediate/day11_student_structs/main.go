package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Grades []int  `json:"grades"`
}

func main() {
	s1 := Student{"Alice", 20, []int{95, 88, 92}}

	fmt.Println(s1)
	res, _ := json.Marshal(s1)

	fmt.Println("encoded json : ",res)

	jsonData, err := json.MarshalIndent(s1, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
