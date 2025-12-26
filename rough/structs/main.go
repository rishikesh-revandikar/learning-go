package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	//initial setup
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// reciever type - attaches func to the struct
func (o *order) changeStatus(status string) {
	o.status = status
}

// * (pointer) needed only when value needs to be updated
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// var order order
	// order := order{
	// 	id:     "1",
	// 	amount: 50.78,
	// 	status: "recieved",
	// }

	myOrder := newOrder("1", 50.67, "recieved")

	myOrder.createdAt = time.Now()

	fmt.Println(*myOrder)
	fmt.Println(myOrder.getAmount())

	myOrder.changeStatus("delieverd")

	fmt.Println(*myOrder)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// newCust := customer{
	// 	name:"Rishi",
	// 	phone: "234234",
	// }

	// newOrder := order{
	// 	id:"2",
	// 	amount: 56.89,
	// 	status: "recieved",
	// 	customer: newCust,
	// }
	newOrder := order{
		id:"2",
		amount: 56.89,
		status: "recieved",
		customer: customer{
			name: "sdsd",
			phone: "432432",
		},
	}

	fmt.Println(newOrder)
}
