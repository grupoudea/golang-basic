package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error al conectar:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "¡Hola, servidor!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error al enviar datos:", err)
		return
	}
	fmt.Printf("Enviado: %s\n", message)
}
