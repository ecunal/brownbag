package main

import "fmt"

func adder(args ...int) {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(adder(1, 2, 3, 4, 5))
}
