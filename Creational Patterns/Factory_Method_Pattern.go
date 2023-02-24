package main

import "fmt"

const (
	french  = 1
	english = 2
)

type translator interface {
	getMessage() string
}

func getTranslator(m int) (translator, error) {
	switch m {
	case french:
		return new(frenchMessage), nil
	case english:
		return new(englishMessage), nil
	default:
		return nil, fmt.Errorf("Unknown building")
	}
}

type frenchMessage string

func (c *frenchMessage) getMessage() string {
	return "Bonjour le monde"
}

type englishMessage string

func (d *englishMessage) getMessage() string {
	return "Hello World"
}

func main() {
	obj1, _ := getTranslator(1)
	fmt.Printf("%T\n", obj1)
	fmt.Println(obj1.getMessage())

	obj2, _ := getTranslator(2)
	fmt.Printf("%T\n", obj2)
	fmt.Println(obj2.getMessage())
}