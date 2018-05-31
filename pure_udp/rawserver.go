package main

import (
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	const proto = (syscall.ETH_P_ALL<<8)&0xff00 | syscall.ETH_P_ALL>>8
	// const proto = (syscall.ETH_P_IP<<8)&0xff00 | syscall.ETH_P_IP>>8
	fmt.Println("\n===== syscall.Socket() =====")
	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_DGRAM, proto)
	// fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, proto)
	defer syscall.Close(fd)

	if_info, _ := net.InterfaceByName("ens4")

	// var haddr [8]byte
	// copy(haddr[0:7], if_info.HardwareAddr[0:7])
	fmt.Println("\n===== syscall.SockaddrLinklayer() =====")
	addr := syscall.SockaddrLinklayer{
		Protocol: proto,
		Ifindex:  if_info.Index,
		// Halen:    uint8(len(if_info.HardwareAddr)),
		// Addr:     haddr,
	}

	fmt.Println("\n===== syscall.Bind() =====")
	if err := syscall.Bind(fd, &addr); err != nil {
		log.Fatal("bind: ", err)
	}

	buf := make([]byte, 65536)
	log.Println("Starting raw server...")
	for {
		fmt.Println("\n===== syscall.Recvfrom() =====")
		n, addr, _ := syscall.Recvfrom(fd, buf, 0)

		fmt.Println("recv byte: ", n)
		sa, _ := addr.(*syscall.SockaddrLinklayer)
		fmt.Printf("Recv SockaddrLinklayer: %+v\n", sa)

	}

}

// func main() {
// 	const proto = (syscall.ETH_P_ALL<<8)&0xff00 | syscall.ETH_P_ALL>>8
// 	fmt.Println("\n===== syscall.Socket() =====")
// 	// fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_IP)
// 	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, proto)
// 	defer syscall.Close(fd)

// 	if_info, _ := net.InterfaceByName("ens4")
// 	var haddr [8]byte
// 	copy(haddr[0:7], if_info.HardwareAddr[0:7])

// 	fmt.Println("\n===== syscall.SockaddrLinklayer() =====")
// 	addr := syscall.SockaddrLinklayer{
// 		Protocol: proto,
// 		Ifindex:  if_info.Index,
// 		Halen:    uint8(len(if_info.HardwareAddr)),
// 		Addr:     haddr,
// 	}

// 	fmt.Println("\n===== syscall.Bind() =====")
// 	if err := syscall.Bind(fd, &addr); err != nil {
// 		log.Fatal("bind: ", err)
// 	}

// 	buf := make([]byte, 65536)
// 	log.Println("Starting raw server...")
// 	for {
// 		fmt.Println("\n===== syscall.Recvfrom() =====")
// 		n, addr, err := syscall.Recvfrom(fd, buf, 0)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		// FOR DEBUG
// 		fmt.Println("recv byte: ", n)
// 		sa, _ := addr.(*syscall.SockaddrLinklayer)
// 		fmt.Printf("Recv SockaddrLinklayer: %+v\n", sa)
// 	}
// }
