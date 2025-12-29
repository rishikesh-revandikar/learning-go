package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Grades []int   `json:"grades"`
	GPA    float64 `json:"gpa"`
}

func (s *Student) calculateGpa() {
	if len(s.Grades) == 0 {
		s.GPA = 0
	}

	totalGrades := 0
	for _, grade := range s.Grades {
		totalGrades += grade
	}

	s.GPA = float64(totalGrades / len(s.Grades))
}

func main() {
	s := Student{
		Name:   "Alice",
		Age:    20,
		Grades: []int{90, 80},
	}

	jsonData, _ := json.MarshalIndent(s, "", "  ")

	fmt.Println(string(jsonData))

	s.calculateGpa()

	newJson, _ := json.MarshalIndent(s, "", "  ")

	fmt.Println(string(newJson))
}
