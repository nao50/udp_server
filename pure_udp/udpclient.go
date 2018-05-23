package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:2152")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("0123456789"))
	if err != nil {
		log.Fatalln(err)
	}
}
