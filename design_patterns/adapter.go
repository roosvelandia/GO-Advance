package main

import "fmt"

// Structural pattern
// adapta una clase para no implementar completamente una interface
// interface with a method
type Payment interface {
	Pay()
}

// example when the method is implemented as original interface
type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

// another instance changing interface implementation
type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using banc account %d  \n", bankAccount)
}

// new object with extra parameters
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// new implementation in the method
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	//bank := &BankPayment{}
	//ProcessPayment(bank)
	// create the adapter object
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	// use the adapter implementation
	ProcessPayment(bpa)
}
