package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
