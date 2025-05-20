package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    float64
	status    string
	customer  Customer // struct nesting or struct embedding
	createdAt time.Time
}
type Customer struct {
	name  string
	email string
}

// kind of constructor
func newOrder(id int, amount float64, status, cname, cemail string) *Order {
	newCust := Customer{
		name:  cname,
		email: cemail,
	}
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
		// customer: Customer{
		// 	name:  cname,
		// 	email: cemail,
		// },
		// -----
		// OR
		/// ----
		customer:  newCust,
		createdAt: time.Now(),
	}
	return &myOrder
}

// available for all
// func changeStatus(order *Order, status string) {
// 	order.status = status
// }

// only for Order struct
func (order *Order) changeStatus(status string) {
	order.status = status
}
func (order Order) getAmount() float64 {
	return order.amount
}
func main() {
	// order1 := Order{
	// 	id:        1,
	// 	amount:    45.90,
	// 	status:    "PENDING",
	// 	createdAt: time.Now(),
	// }
	// // var order2 Order = Order{1, 3.14, "pending", time.Now()}
	// // order1.amount = 60.9

	// fmt.Println(order1)
	// order1.changeStatus("DONE")
	// fmt.Println(order1.getAmount())
	od1 := newOrder(1, 3.14, "DONE", "utsav", "utsav@gmail.com")
	od1.customer.name = "Kevin"
	fmt.Println(od1)

	// inline struct
	// languages := struct {
	// 	name   string
	// 	isGood bool
	// }{"go lang", true}

	// fmt.Println(languages.name)
}
