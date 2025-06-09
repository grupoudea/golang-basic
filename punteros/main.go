package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type VertexV2 struct {
	X, Y int
}

var (
	v1 = VertexV2{1, 2}  // tiene tipo Vertex
	v2 = VertexV2{X: 1}  // Y:0 es implicito
	v3 = VertexV2{}      // X:0 y Y:0
	p1 = &VertexV2{1, 2} // tiene tipo *Vertex
)

func main() {

	diccionario := map[string]Vertex{}

	diccionario["1"] = Vertex{222, 22}
	key := diccionario["1"]
	fmt.Println("key %d value %d", key.X, key.Y)

	for key, value := range diccionario {
		fmt.Println("Printing...")
		fmt.Println(key, value)
	}

	v := Vertex{1, 2}
	p := &v
	p.X = 4
	fmt.Println(v)
	fmt.Println(v1, p1, v2, v3)

}
