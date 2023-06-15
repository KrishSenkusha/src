package main

import "fmt"

var name string
var age int

func print(n string, a int) {
	var newage string = fmt.Sprint(a)
	fmt.Println("Your name is: " + n)
	fmt.Println("Your age is: " + newage)
}

func main() {
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	print(name, age)
}
