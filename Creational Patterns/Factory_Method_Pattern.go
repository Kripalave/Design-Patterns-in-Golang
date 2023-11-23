package main

import "fmt"

// Product: Payment interface
type Payment interface {
	Pay(amount float64) string
}

// ConcreteProductA: Credit Card Payment
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card", amount)
}

// ConcreteProductB: PayPal Payment
type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

// Creator: PaymentFactory interface
type PaymentFactory interface {
	CreatePayment() Payment
}

// ConcreteCreatorA: CreditCardPaymentFactory
type CreditCardPaymentFactory struct{}

func (c *CreditCardPaymentFactory) CreatePayment() Payment {
	return &CreditCardPayment{}
}

// ConcreteCreatorB: PayPalPaymentFactory
type PayPalPaymentFactory struct{}

func (p *PayPalPaymentFactory) CreatePayment() Payment {
	return &PayPalPayment{}
}

func main() {
	// Client code
	creditCardFactory := &CreditCardPaymentFactory{}
	creditCardPayment := creditCardFactory.CreatePayment()
	fmt.Println(creditCardPayment.Pay(100.0))

	payPalFactory := &PayPalPaymentFactory{}
	payPalPayment := payPalFactory.CreatePayment()
	fmt.Println(payPalPayment.Pay(50.0))
}
