package main

import "fmt"

type OrderStatus int

// type OrderStatus string

const (
	Recevied OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

// const (
// 	Recevied  OrderStatus = "recevied"
// 	Confirmed             = "confirmed"
// 	Prepared              = "prepared"
// 	Delivered             = "delivered"
// )

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Change Order to ", status)
}
func main() {
	changeOrderStatus(Recevied)
	// changeOrderStatus("recevied")
}
