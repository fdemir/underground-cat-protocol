package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// If the purpose is client, then run the client function, for test
	if len(os.Args) > 1 && os.Args[1] == "c" {
		Client()
		return
	}

	fmt.Println("Starting UCP server on :8080")
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go HandleConnection(conn)
	}
}
