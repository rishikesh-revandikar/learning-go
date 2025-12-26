package main

import "fmt"

type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32){
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Making payment using razorpay",amount)
}

func (r razorpay) refund(amount float32, account string){
	// reufnd logic
}

type stripe struct{}

func (r stripe) pay(amount float32){
	// logic to make payment
	fmt.Println("Making payment using stripe",amount)
}

func (r stripe) refund(amount float32, account string){
	// reufnd logic
}
func main(){
	// razorPayGw := razorpay{}
	stripeGw := stripe{}
	newPayment := payment{
		gateway: stripeGw,
	}
	newPayment.makePayment(100)
}