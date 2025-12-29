package main

import (
	"fmt"
	"sync"
)

func Runner(num int, wg *sync.WaitGroup){
	defer func(){wg.Done()}()
	fmt.Println("Runner ",num)
}

func main()  {
	var wg sync.WaitGroup

	for i:=0; i < 5; i++{
		wg.Add(1)
		go Runner(i,&wg)
	}

	wg.Wait()


}
