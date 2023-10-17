package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())

	s := "ss"
	integer, error := strconv.Atoi(s)

	if error != nil {
		fmt.Println("Error: ", error)
	}

	n := 42
	s1 := strconv.Itoa(n)

	fmt.Printf("valor del entero es %d y el string es %s\n", integer, s1)

}
