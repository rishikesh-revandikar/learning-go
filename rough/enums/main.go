package main

import "fmt"

// enumerated types

type OrderStatus string

const (
	Recieved  OrderStatus = "recieved"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ", status)
}

func main() {
	changeOrderStatus(Recieved)
}
