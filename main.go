package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	fmt.Println("It all starts here")
	connection, err := icmp.ListenPacket("udp4", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	byteMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 8,
		Body: &icmp.Echo{
			ID:   os.Getpid(),
			Seq:  1,
			Data: []byte("hello test"),
		},
	}

	message, err := byteMessage.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := connection.WriteTo(message, &net.IPAddr{IP: net.ParseIP("127.0.0.1"), Zone: "en0"}); err != nil {
		log.Fatal(err)
	}

	rb := make([]byte, 1500)
	n, _, err := connection.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}

	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got %+v", rm)

}
