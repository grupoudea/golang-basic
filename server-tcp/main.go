package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error al escuchar:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Servidor escuchando en localhost:9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Conexión desde cliente %v \n", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer datos:", err)
	}
	fmt.Printf("Recibido: %s\n", string(buffer))
	conn.Close()
}
