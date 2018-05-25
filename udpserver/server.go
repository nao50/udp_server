package main

import "sync"

// "github.com/naoyamaguchi/udp_server/tbf"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go uplink()
	go downlink()

	wg.Wait()
}

// func main() {
// 	udpAddr := &net.UDPAddr{
// 		IP:   net.ParseIP("127.0.0.1"),
// 		Port: 2152,
// 	}
// 	updLn, err := net.ListenUDP("udp", udpAddr)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Tocken Bucket Filter
// 	// // TODO(nao): 接続ごとに値を変えたい
// 	ctx, limit := tbf.InitTokenBucket()

// 	// Buffer
// 	buf := make([]byte, 1024)
// 	log.Println("Starting udp server...")

// 	for {
// 		n, addr, err := updLn.ReadFromUDP(buf)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		start := time.Now()

// 		err = tbf.TokenBucketFilter(ctx, n, limit)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		go func() {
// 			log.Println("size: ", n)
// 			log.Printf("Reciving data: %s from %s", string(buf[:n]), addr.String())
// 			// updLn.WriteTo(buf[:n], sendUDPAddr)
// 			updLn.WriteTo(buf[:n], addr)
// 		}()
// 		log.Println("End test. ", time.Since(start))
// 	}
// }
func uplink() {
	uplinkBuffer := make([]byte, 1550)
	ctx, limit := tbf.InitTokenBucket()
	///////////////////////////////////////////////////////////////	
	sendUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("10.0.11.10"),
		Port: 2153,
	}
	// sendUdpConn, err := net.ListenUDP("udp", sendUdpAddr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	///////////////////////////////////////////////////////////////
	recvUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 2152,
	}
	recvUdpConn, err := net.ListenUDP("udp", recvUdpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	///////////////////////////////////////////////////////////////
	for {
		n, addr, err := recvUdpConn.ReadFromUDP(uplinkBuffer)
		if err != nil {
			log.Fatalln(err)
		}

		// tbf
		err = tbf.TokenBucketFilter(ctx, n, limit)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			// recvUdpConn.WriteTo(buf[:n], addr)
			recvUdpConn.WriteTo(uplinkBuffer[:n], sendUdpAddr)
			// sendUdpConn.WriteTo(uplinkBuffer[:n], sendUdpAddr)
		}()

	}

}



func downlink() {
	uplinkBuffer := make([]byte, 1550)
	ctx, limit := tbf.InitTokenBucket()
	///////////////////////////////////////////////////////////////	
	sendUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("10.0.10.10"),
		Port: 2253,
	}
	// sendUdpConn, err := net.ListenUDP("udp", sendUdpAddr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	///////////////////////////////////////////////////////////////
	recvUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 2252,
	}
	recvUdpConn, err := net.ListenUDP("udp", recvUdpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	///////////////////////////////////////////////////////////////
	for {
		n, addr, err := recvUdpConn.ReadFromUDP(uplinkBuffer)
		if err != nil {
			log.Fatalln(err)
		}

		// tbf
		err = tbf.TokenBucketFilter(ctx, n, limit)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			// recvUdpConn.WriteTo(buf[:n], addr)
			recvUdpConn.WriteTo(uplinkBuffer[:n], sendUdpAddr)
			// sendUdpConn.WriteTo(uplinkBuffer[:n], sendUdpAddr)
		}()

	}

}

