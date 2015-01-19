package main

import (
	"errors"
	"fmt"
)

func foo(i int) (str string, err error) {

	str = "return"

	defer func() { str += " string" }()

	if i < 0 {
		err = errors.New("Acayip error")
		return
	}
	// code continued
	return
}

func main() {
	if str, err := foo(5); err != nil {
		panic("Error in foo!")
	} else {
		fmt.Println(str)
	}
}

// END OMIT
