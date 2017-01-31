package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Jullian":
		fmt.Println("Wassup Jullian")
	case "Sushant":
		fmt.Println("Wassup Susant")
	default:
		fmt.Println("You have no friends")
	}
}
