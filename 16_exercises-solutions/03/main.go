package main

import "fmt"

func greatest(n ...int) int {
	var g int
	for _, num := range n {
		if num > g {
			g = num
		}
	}
	return g
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4, 5))
}
