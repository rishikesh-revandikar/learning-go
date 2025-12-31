package main

import (
	"fmt"
	"sync"
	"time"
)

func api_call(api int, wg *sync.WaitGroup){
	defer func(){
		wg.Done()
		fmt.Println("API req completed :: ",api)
	}()
	fmt.Println("API req in processing :: ",api)
	time.Sleep(time.Second * 5)

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++{
		wg.Add(1)
		go api_call(i,&wg)
	}
	wg.Wait()

	fmt.Println("API calls completed!!")
}
