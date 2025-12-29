package main

import (
	"fmt"
	"sync"
)


func main(){
	ping := make(chan string)
	pong := make(chan string)


	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i:=0; i < 5; i++{
			msgSent := fmt.Sprintf("Ping %v",i)
			ping <- msgSent

			msgRecv := <- pong
			fmt.Printf("Ping recieved : %s\n",msgRecv)
		}
	}()

	go func(){
		defer wg.Done()
		for i:=0; i < 5; i++{
			msgRecv := <- ping
			fmt.Printf("Pong recieved : %s\n",msgRecv)

			msgSent := fmt.Sprintf("Pong %v",i)
			pong <- msgSent
		}
	}()

	wg.Wait()

	close(ping)
	close(pong)

	fmt.Println("Conversation finished and channels closed.")
}