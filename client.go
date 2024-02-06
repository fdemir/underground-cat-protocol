package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Say something: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server said: " + message)
	}
}
