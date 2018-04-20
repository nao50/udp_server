package main

import (
	"context"
	"log"
	"net"
	"time"

	"golang.org/x/time/rate"
)

const (
	M = 8000 // 1秒あたりの処理制限
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	updLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
	}

	// Tocken Bucket Filter
	ctx := context.Background()
	n := rate.Every(time.Second / M)
	l := rate.NewLimiter(n, M) //必ずしも上限Mである必要はない。ここの上限値がバースト値となる。

	// Buffer
	buf := make([]byte, 1024)
	log.Println("Starting udp server...")

	for {
		n, addr, err := updLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}

		start := time.Now()

		// TBF n[byte]分のTockenをマイナス。bpsはbit/secであることに注意が必要。
		if err := l.WaitN(ctx, n); err != nil {
			log.Fatalln(err)
		}

		go func() {
			log.Println(n)
			log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())
			updLn.WriteTo(buf[:n], addr)
		}()
		log.Println("End test. ", time.Since(start))
	}
}
