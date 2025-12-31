package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct{
	ID int
	Name string
}


func main(){
	fileName := "data.txt"

	user := User{
		ID : 1,
		Name: "Rishikesh",
	}

	file, err := os.Create(fileName)
	if err != nil{
		fmt.Println("Error creating file : ",err)
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%d,%s\n", user.ID, user.Name)

	if err!=nil{
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data successfully written to data.txt")

	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	fmt.Println("Reading data back:")

	for scanner.Scan() {
		line := scanner.Text()
		
		// Parse the data back into a struct format
		var id int
		var name string
		fmt.Sscanf(line, "%d,%s", &id, &name)
		
		fmt.Printf("Parsed Struct -> ID: %d, Name: %s\n", id, name)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}

}