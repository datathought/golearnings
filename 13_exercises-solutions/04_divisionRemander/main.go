package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int
	fmt.Printf("Enter a Small Number: ")
	fmt.Scan(&smallNumber)
	fmt.Printf("Enter a Large Number: ")
	fmt.Scan(&largeNumber)
	remainder := largeNumber % smallNumber
	fmt.Println("Remander = ", remainder)
}
