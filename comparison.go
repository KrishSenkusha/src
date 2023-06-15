package main

import "fmt"

func main() {
	a := 21
	b := 42

	fmt.Println("a==b: ", a == b)   // false
	fmt.Println("a==b: ", a == a)   // true
	fmt.Println("a != b: ", a != b) // true
	fmt.Println("a>b: ", a > b)     // false
	fmt.Println("a<b: ", a < b)     // true
	fmt.Println("a<=b: ", a <= b)   // true
	fmt.Println("a<=a: ", a <= a)   // true
	fmt.Println("a>=b: ", a >= a)   // true
}
