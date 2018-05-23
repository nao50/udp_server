package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2152,
	}
	fmt.Println("\n===== net.ListenUDP() =====")
	updConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 1550)
	for {
		fmt.Println("\n===== updConn.ReadFromUDP() =====")
		n, raddr, err := updConn.ReadFromUDP(buf)

		fmt.Println("recieved size: ", n)
		fmt.Println("remote address: ", raddr)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
