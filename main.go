package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b	
}

func main() {
	number := 25
	age := 32

	fmt.Println("The number is: ", number)
	fmt.Println("My age is: ", age)
	fmt.Println("The sum 6 + 6 is:", sum(6, 6))
	fmt.Println("The multi 6 * 2 is:", multi(6, 2))
}
