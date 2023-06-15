package main

import "fmt"

func main() {
	var floor int
	var pass int
	fmt.Println("Enter floor and password: ")
	fmt.Scanln(&floor)
	fmt.Scanln(&pass)
	fmt.Println(floor, pass)

	var name string
	fmt.Scanln(&name)
	fmt.Println(name)

	number := 12
	fmt.Println(number)

	const a int = 13
	fmt.Println(a)
}
