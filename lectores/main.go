package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Agregar un metodo Read([]byte) (int, error) a MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil

}

func main() {
	reader.Validate(MyReader{})
}
