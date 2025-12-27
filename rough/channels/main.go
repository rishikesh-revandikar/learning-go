package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("Processing number ", <-numChan)
}

func sum(result chan int, num1 int, num2 int) {
	res := num1 + num2
	result <- res
}

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("processing...")

}

func emailSender(emailChan chan string, done chan bool){
	defer func() {
		done <- true
	}()
	for email := range emailChan{
		fmt.Println("Sending email to ",email)
		time.Sleep(time.Second)
	}
}

func main() {

	// time.Sleep(time.Second * 2)
	// messageChan := make(chan string)

	// // channels are blocking

	// messageChan <- "Heloo"
	// msg := <- messageChan

	// fmt.Println(msg)

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result // blocking

	// fmt.Println(res)

	// done := make(chan bool)

	// go task(done)

	// <-done

	// Buffered channel -> can send limited amount of data without blocking unlike traditional channels
	//                     where only one data is sent and is blocking execution


	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan,done)

	for i:=0; i<100; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("Done sending...")
	// emailChan <- "1@ex.com"
	// emailChan <- "2@ex.com"

	// fmt.Println(<- emailChan)
	// fmt.Println(<- emailChan)

	<- done
}
