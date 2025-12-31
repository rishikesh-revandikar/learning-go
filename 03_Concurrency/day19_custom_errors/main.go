package main

import (
	"errors"
	"fmt"
)

type MyError struct{
	Code int 
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s\n",e.Code,e.Message)
}

var ErrNotFound = &MyError{
	Code: 404,
	Message: "Not Found",
}

func databaseQuery() error {
	return fmt.Errorf("database operation failed: %w",ErrNotFound)
}

func main()  {
	err := databaseQuery()

	if errors.Is(err,ErrNotFound){
		fmt.Println("Caught a specific 404 error!")
	} else {
		fmt.Println("Unknown error occurred.")
	}
	
	fmt.Println("Original Error:", err)
}