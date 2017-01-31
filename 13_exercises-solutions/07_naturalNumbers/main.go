package main

import "fmt"

func main() {
	var naturalNumbers int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			naturalNumbers += i

		} else if i%5 == 0 {
			naturalNumbers += i

		}
	}
	fmt.Println("sum is ", naturalNumbers)
}
