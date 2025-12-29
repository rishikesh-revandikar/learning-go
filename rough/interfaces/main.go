package main

import (
	"fmt"
	"math"
	"time"
)

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


type Abser interface{
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct{
	x,y float64
}

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vertex) String() string{
	return fmt.Sprintf("(x: %v, y: %v)",v.x,v.y)
}

type MyError struct{
	when time.Time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.when, e.What)
}


func run() error{
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}


func main(){
	// razorPayGw := razorpay{}
	// stripeGw := stripe{}
	// newPayment := payment{
	// 	gateway: stripeGw,
	// }
	// newPayment.makePayment(100)

	// var a Abser

	// f := Myfloat(-math.Sqrt2)
	// v := Vertex{3,4}

	// a = f
	// a = &v

	// // a = v

	// fmt.Println(a)

	if err := run(); err != nil{
		fmt.Println(err)
	}
}