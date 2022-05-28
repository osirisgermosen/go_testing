package main

import "fmt"

func sum(a, b int) {
	return a + b
}

func main() {
	number := 25
	age := 32
	fmt.Println("The number is: ", number)
	fmt.Println("My age is: ", age)
	fmt.Println("The sum 5 + 6 is:", sum(5, 6))
}
