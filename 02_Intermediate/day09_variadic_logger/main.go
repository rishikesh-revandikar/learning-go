package main

import (
	"fmt"
	"strings"
)


func Log(prefix string, messages ...string){
	defer fmt.Printf("[%s] --- Execution Finished ---\n", prefix)

	fullmsg := strings.Join(messages," ")

	fmt.Printf("[%s] %s\n", prefix, fullmsg)
}


func main()  {
	Log("INFO", "Starting", "the", "system...")
	fmt.Println()
	Log("ERROR", "Database", "connection", "failed!")
}