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
	// p1 := Payment{}
	// p1.makeAPayment(100)

	// p2 := Payment{}
	// p2.makeAPayment(20)
	// stripeGateway := Stripe{}
	// razorpayGatway := RazorPay{}

	// fGatway := fakePayemnt{}
	paypalGatway := PayPal{}
	p1 := Payment{
		gatWay: paypalGatway,
	}
	p1.makePayment(200)
}
