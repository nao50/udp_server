package main

import (
	"log"
	"net"
	// "github.com/naoyamaguchi/udp_server/tbf"
	"./tbf"
)

func downlink() {
	uplinkBuffer := make([]byte, 1550)
	ctx, limit := tbf.InitTokenBucket()
	///////////////////////////////////////////////////////////////	
	sendUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("10.0.10.10"),
		Port: 2253,
	}
	sendUdpConn, err := net.ListenUDP("udp", sendUdpAddr)
	if err != nil {
		log.Fatalln(err)
	}
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
			recvUdpConn.WriteTo(buf[:n], sendUdpAddr)
			sendUdpConn.WriteTo(buf[:n], sendUdpAddr)
		}()

	}

}
