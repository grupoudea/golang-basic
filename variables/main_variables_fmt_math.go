package main

import (
	"fmt"
	"math"
)

const PI = 3.14

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Feliz día de", PI)

	a1, b := swap("hola", "mundo")
	fmt.Println(a1, b)

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
