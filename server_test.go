package main

import (
	"log"
	"net"
	"os"
	"testing"
)

func BenchmarkUdp(b *testing.B) {
	// セッション使い回し
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		/*
					conn, err := net.Dial("udp", "127.0.0.1:8080")
					if err != nil {
						log.Fatalln(err)
						os.Exit(1)
					}
			    defer conn.Close()
		*/
		_, err = conn.Write([]byte("0123456789"))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		recvBuf := make([]byte, 1024)
		n, err := conn.Read(recvBuf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		log.Printf("Received data: %s", string(recvBuf[:n]))
		//conn.Close()

	}
}
