package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
	fmt.Fprintf(connection, "E ai servidor")
	_, _, err = connection.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", buffer)
}

func clientLog() {
	f, _ := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()

	log.SetOutput(f)
	//log.Println("njasjnxja")
}
