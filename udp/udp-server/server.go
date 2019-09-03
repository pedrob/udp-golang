package main

import (
	"fmt"
	"net"
)

//without check errors

func main() {
	listenClients()
}

func listenClients() {
	buffer := make([]byte, 1024)
	address, _ := net.ResolveUDPAddr("udp", "localhost:8888")
	server, _ := net.ListenUDP("udp", address)
	for {
		_, clientAddr, _ := server.ReadFromUDP(buffer)
		fmt.Printf("%s", buffer)
		go sendToClient(server, clientAddr)
	}
}

func sendToClient(connection *net.UDPConn, clientAddr *net.UDPAddr) {
	connection.WriteToUDP([]byte("E ai broder"), clientAddr)
}
