package main

import (
	"log"
	"net"
	"time"

	// "github.com/naoyamaguchi/udp_server/tbf"
	"./tbf"
)

func main() {
	// sendUDPAddr := &net.UDPAddr{
	// 	IP:   net.ParseIP("127.0.0.1"),
	// 	Port: 2121,
	// }

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2152,
	}
	updLn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Tocken Bucket Filter
	// // TODO(nao): 接続ごとに値を変えたい
	ctx, limit := tbf.InitTokenBucket()

	// Buffer
	buf := make([]byte, 1024)
	log.Println("Starting udp server...")

	for {
		n, addr, err := updLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}

		start := time.Now()

		err = tbf.TokenBucketFilter(ctx, n, limit)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			log.Println("size: ", n)
			log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())
			// updLn.WriteTo(buf[:n], sendUDPAddr)
			updLn.WriteTo(buf[:n], addr)
		}()
		log.Println("End test. ", time.Since(start))
	}
}
