package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("\n===== net.Listen() =====")
	ln, err := net.Listen("udp", ":2152")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// https://golang.org/pkg/net/#Listen
		// func Listen is not support udp??
		fmt.Println("\n===== ln.Accept() =====")
		conn, err := ln.Addr
		// if err != nil {
		// 	log.Print(err)
		// 	continue
		// }

	}
}
