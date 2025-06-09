package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ti := time.Now()
	for i := 0; i < 100; i++ {
		fmt.Println("iteración %d", (i + 1))
	}
	tf := time.Now()
	fmt.Println("Tiempo transcurrido: %.2f", tf.Sub(ti).Seconds())
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
