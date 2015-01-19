package main

import "fmt"

type MyInt int

func (m MyInt) Abs() int {
	i := int(m)
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	i := MyInt(-10)
	fmt.Println(i.Abs())
}

// END OMIT
