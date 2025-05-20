- [Struct](#struct)
- [interface](#interface)
- [enums](#enums)
- [generics](#generics)

## Struct
- Struct is s user define data types that grouped different type of variable under the single name , in go their are no classes .

- struct declaration 
```go
type Order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
}
```
```go
order1 := Order{
	id:        1,
	amount:    45.90,
	status:    "PENDING",
	createdAt: time.Now(),
}
var order2 Order = Order{1, 3.14, "pending", time.Now()}

order1.amount = 60.9 // modification by dot
fmt.Println(order1) // print entire order
```

Methods for struct

```go
func (order *Order) changeStatus(status string) {
	order.status = status // no need of Dereferences
}
func (order Order) getAmount() float64 {
	return order.amount 
}

// kind of constructor
func newOrder(id int, amount float64, status string) *Order { // Usually return via pointer
	myOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}
```
---
```go
languages := struct {
	name   string
	isGood bool
}{"go lang", true}

```
## interface
- interface are just a template
- It define the set of methods signature but does not provide implementations 
- Go uses implicit implementation, meaning a type does not need to explicitly declare that it implements an interface—it just needs to define the required methods.

```go
package main

import "fmt"

type Payment struct {
}

func (p Payment) makeAPayment(amount float64) {
	// razorPayInstance := RazorPay{}
	// razorPayInstance.pay(amount)

	stripeGatway := Stripe{}
	stripeGatway.pay(123)
}

type RazorPay struct {
}

func (r RazorPay) pay(amount float64) {
	fmt.Println("Making payment via RazorPay amount of", amount)
}

type Stripe struct {
}

func (r Stripe) pay(amount float64) {
	fmt.Println("Making payment via Stripe amount of", amount)
}

func main() {
	p1 := Payment{}
	p1.makeAPayment(100)

	p2 := Payment{}
	p2.makeAPayment(20)
}
```
- here first i complete payment via RazorPay but latter on i required to complete payment via stripe , so i have to comment some code in payment struct and need to create Stripe struct hard coding of value
```go
package main

import "fmt"

type Payment struct {
	gatWay Stripe
	// gatWay RazorPay
}

func (p Payment) makeAPayment(amount float64) {
	p.gatWay.pay(amount)
}

type RazorPay struct {
}

func (r RazorPay) pay(amount float64) {
	fmt.Println("Making payment via RazorPay amount of", amount)
}

type Stripe struct {
}

func (r Stripe) pay(amount float64) {
	fmt.Println("Making payment via Stripe amount of", amount)
}

func main() {
	stripeGateway := Stripe{}
	p1 := Payment{
		gatWay: stripeGateway,
	}
	p1.makeAPayment(200)
}

```
- slightly better but still hard coded due paymentGateway is hard coded i can't pass RazorPay  
- via interface

```go
package main

import "fmt"

type Paymenter interface {
	pay(amount float64)
	refund(amount float64, accountNum int)
}

type Payment struct {
	gatWay Paymenter
}

func (p Payment) makePayment(amount float64) {
	p.gatWay.pay(amount)
}

type RazorPay struct{}

func (r RazorPay) pay(amount float64) {
	fmt.Println("Making payment via RazorPay amount of", amount)
}

type Stripe struct{}

func (r Stripe) pay(amount float64) {
	fmt.Println("Making payment via Stripe amount of", amount)
}

type fakePayemnt struct{}

func (p fakePayemnt) pay(amount float64) {
	fmt.Println("Making payment via fake Payemnt for testing amount of", amount)
}

type PayPal struct{}

func (p PayPal) pay(amount float64) {
	fmt.Println("Making payment via paypal amount of", amount)
}
func (p PayPal) refund(amount float64, accountNum int) {
	fmt.Printf("Amount %f refung for account %d", amount, accountNum)
}

func main() {
	paypalGatway := PayPal{}
	p1 := Payment{
		gatWay: paypalGatway,
	}
	p1.makePayment(200)
}

```
- after using interface , PayPal gateway were added without need of any code modification , just by writing more code 
- just remember that struct should implements all methods that is define in interface and method signature should be same
## enums
## generics