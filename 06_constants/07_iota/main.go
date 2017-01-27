package main

import "fmt"

const (
	_ = iota //0
	//KB this is bit wise shifting 2^(1*10) = 1024 or in binary 1000000000
	KB = 1 << (iota * 10)

	//MB this is bit wise shifting 2^(2*10) = 1048576 or in binary 100000000000000000000
	MB = 1 << (iota * 10)

	//GB this is bit wise shifting 2^(3*10) = 1073741824 or in binary 1000000000000000000000000000000
	GB = 1 << (iota * 10)

	//TB this is bit wise shifting 2^(3*10) = 1099511627776 or in binary 10000000000000000000000000000000000000000
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
