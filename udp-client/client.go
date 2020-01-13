package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	buffer := make([]byte, 1024)
	addr, err := net.ResolveUDPAddr("udp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
	}
	connection, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
	defer connection.Close()
	//envia os dados
	fmt.Fprintf(connection, "E ai servidor")
	err = connection.SetDeadline(time.Now().Add(time.Second * 15))
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
	n, server, err := connection.ReadFromUDP(buffer)
	fmt.Printf("msg: %s\nadress: %s\nmsg length: %d\nmsg time: %s\n", buffer, server, n, time.Now().Local())
}
