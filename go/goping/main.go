package main

import (
	"log"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const icmpProtocolNumber = 1

func main() {
	ping("google.com")
}

func ping(addr string) {
	dstIP, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		panic(err)
	}

	udpServer, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()

	icmpMsg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	icmpMsgBs, err := icmpMsg.Marshal(nil)
	if err != nil {
		panic(err)
	}

	if _, err := udpServer.WriteTo(icmpMsgBs, dstIP); err != nil {
		panic(err)
	}

	buffer := make([]byte, 1500)
	bsRead, peer, err := udpServer.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}

	recvIcmpMsg, err := icmp.ParseMessage(icmpProtocolNumber, buffer[:bsRead])
	if err != nil {
		log.Fatal(err)
	}
	switch recvIcmpMsg.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %s (%v)", addr, peer)
	default:
		log.Printf("got %+v; want echo reply", recvIcmpMsg)
	}
}
