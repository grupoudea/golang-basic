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

	//LEER DEL SERVER
	buffer := make([]byte, 1024)
	_, err2 := conn.Read(buffer)
	if err2 != nil {
		fmt.Println("Error al leer datos:", err)
	}
	fmt.Printf("Recibido: %s\n", string(buffer))

}
