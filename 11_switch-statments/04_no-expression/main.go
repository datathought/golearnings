package main

import "fmt"

func main() {

	myFriendsName := "Al"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friends of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Wassup My M friends")
	case myFriendsName == "Jullian":
		fmt.Println("Wassup Jullian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Susant")
	default:
		fmt.Println("You have no friends")
	}
}
