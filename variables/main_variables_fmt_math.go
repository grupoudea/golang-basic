package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, o float32

	fmt.Print("Ingrese el valor del lado opuesto: ")
	fmt.Scanln(&o)

	fmt.Print("Ingrese el valor del lado adyacente: ")
	fmt.Scanln(&a)

	fmt.Printf("o:%g  a:%g \n", o, a)

	h = float32(math.Sqrt(math.Pow(float64(o), 2) + math.Pow(float64(a), 2)))

	fmt.Println("Hipotenusa: ", h)

	area_triangulo := (a * o) / 2

	fmt.Printf("Area: %.2f\n", area_triangulo)

	perimetro := (h + a + o)

	fmt.Printf("Perimetro: %.2f\n", perimetro)

}
