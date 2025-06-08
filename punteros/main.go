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
	v := Vertex{1, 2}
	p := &v
	p.X = 4
	fmt.Println(v)
	fmt.Println(v1, p1, v2, v3)

}
