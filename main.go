package main

import (
	"fmt"
	"time"
)

func say(s string){
	for i:=0; i < 5; i++{
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}


func main() {
	// go say("world")
	// say("hello")

	tick := time.Tick(100*time.Millisecond)
	boom := time.After(500*time.Millisecond)

	for{
		select{
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
