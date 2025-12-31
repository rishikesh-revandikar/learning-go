package main

import (
	"fmt"
	"sync"
	"time"
)

func api_call(api int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Println("API req completed :: ", api)
	}()
	fmt.Println("API req in processing :: ", api)
	time.Sleep(time.Second * 5)

}

func api_call_chan(req chan int, done chan bool) {
	for {
		select {
		case api := <-req:
			fmt.Println("API req in processing :: ", api)
			time.Sleep(time.Second * 5)
			fmt.Println("API req completed :: ", api)
		default:
			done<-true
		}
	}

}

func main() {

	/*
	* Waitgroup implementation
	 */

	// var wg sync.WaitGroup

	// for i := 0; i < 3; i++{
	// 	wg.Add(1)
	// 	go api_call(i,&wg)
	// }
	// wg.Wait()

	requests := make(chan int, 10)
	done := make(chan bool)

	for i := 0; i < 5; i++ {
		requests <- (i + 1)
	}

	go api_call_chan(requests,done)

	<-done
	fmt.Println("API calls completed!!")
}
