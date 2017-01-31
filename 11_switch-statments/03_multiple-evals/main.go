package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both names start with M")
	case "Jullian", "Sushant":
		fmt.Println("Wassup Jullian / Susant")
	default:
		fmt.Println("You have no friends")
	}
}
