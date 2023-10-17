package main

import (
	"fmt"
	"runtime"
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

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Estás en windows")
	case "linux":
		fmt.Println("Estás en Linux")
	case "darwin":
		fmt.Println("Estás en Mac OS X")
	default:
		fmt.Println("Estás en otro")
	}

}
