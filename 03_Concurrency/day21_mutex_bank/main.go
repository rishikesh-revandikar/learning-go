package main

import (
	"fmt"
	"sync"
)

type Account struct{
	amount int
	m sync.Mutex
}

func (a *Account) deposit(){
	a.m.Lock()
	defer a.m.Unlock()
	a.amount += 1
}

func main(){

	newAcc := &Account{}

	var wg sync.WaitGroup

	wg.Add(100)

	for i:=0; i < 100 ; i++{
		go func(){
			defer wg.Done()
			newAcc.deposit()
		}()
	}

	wg.Wait()

	fmt.Println(newAcc.amount)
}