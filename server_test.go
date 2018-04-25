package main

import (
	"log"
	"net"
	"os"
	"testing"
)

func BenchmarkUdp(b *testing.B) {
	// セッション使い回し
	conn, err := net.Dial("udp", "127.0.0.1:2152")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = conn.Write([]byte("0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		// recvBuf := make([]byte, 1024)
		// n, err := conn.Read(recvBuf)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	os.Exit(1)
		// }

		// log.Printf("Received data: %s", string(recvBuf[:n]))

	}
}

// 	M = 80000000
// 10byte -> 20000	     74795 ns/op	    1104 B/op	       5 allocs/op

// 	M = 80000000
// 100byte -> 20000	     71912 ns/op	    1392 B/op	       5 allocs/op

// 	M = 80000000
//  send only
// 100byte -> 300000	      4392 ns/op	     112 B/op	       1 allocs/op
// 100byte / 4 microsec -> (100 / 1000000 [Mbyte])  /  (4 * 1000000 [s]) = 25[Mbyte/s] = 200Mbit/s
