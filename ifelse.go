package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter Value of a: ")
	fmt.Scanln(&a)
	fmt.Println("Enter value of b: ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("a+b: ", a+b)
	} else {
		fmt.Println("a*b:", a*b)
	}
}
