package main

import "fmt"

func prime(nth int) int {
	//n is the number to be tested for prime
	//p is the number of prime numbers found
	var n, p int
	for {
		n++
		//i is the divsor to find a prime number
		for i := 1; n >= i; i++ {
			//i is tested for being the number 1 because prime numbers are alowed to be diveded by 1
			if i == 1 {
				continue
			}
			//if no number is found by the time divsor n equels the number i being tested then its prime
			if i == n {
				p++
			}
			//breaks out of loop and starts test for next number as the current number is not prime
			if n%i == 0 {
				break
			}
		}
		//Testing to see if the current prime number is the prime number at the nth.
		if p == nth {
			//fmt.Printf("The %v prime number is: %v \n", nth, n)
			return n
		}
	}

}

func main() {
	fmt.Println(prime(10001))
}

/*
10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/
