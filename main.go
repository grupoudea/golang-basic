package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Mañana")
	} else if hora < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}

	fmt.Println("Time: ", t)

}
