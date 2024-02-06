package main

import (
	"bufio"
	"net"
	"strings"
)

const (
	meow     = "meow"
	scratch  = "scratch"
	purr     = "purr"
	hiss     = "hiss"
	sleep    = "sleep"
	wake     = "wake"
	meowmeow = "meowmeow"
	zz       = "zz"
	yawn     = "yawn"
	hrr      = "hrr"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		catm := strings.TrimSpace(msg)

		switch catm {
		case meow:
			conn.Write([]byte(meowmeow + "\n"))
		case scratch:
			conn.Write([]byte(scratch + "\n"))
		case purr:
			conn.Write([]byte(hiss + "\n"))
		case sleep:
			conn.Write([]byte(zz + "\n"))
		case wake:
			conn.Write([]byte(yawn + "\n"))
		default:
			conn.Write([]byte(hrr + "\n"))
		}
	}
}
