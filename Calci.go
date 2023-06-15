package main

import "fmt"

func main() {
	var a, b, result float64
	var ch string
	fmt.Println("Enter Number 1: ")
	fmt.Scanln(&a)

	fmt.Println("Enter Number 2: ")
	fmt.Scanln(&b)

	fmt.Println("Enter operation you want to perform: ")
	fmt.Scanln(&ch)

	switch ch {
	case "+":
		result = a + b
		fmt.Println("Result:", result)
	case "-":
		result = a - b
		fmt.Println("Result:", result)
	case "*":
		result = a * b
		fmt.Println("Result:", result)
	case "/":
		result = a / b
		fmt.Println("Result:", result)
	default:
		fmt.Println("Invalid Input")
	}
}
