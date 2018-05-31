package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("\n===== net.Dial() =====")
	conn, err := net.Dial("udp", "127.0.0.1:2152")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Println("\n===== conn.Write() =====")
	_, err = conn.Write([]byte("0123456789"))
	if err != nil {
		log.Fatalln(err)
	}
}
