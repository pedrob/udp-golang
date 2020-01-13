package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	listenClients()
}

var (
	receivedPackets = 0
	sentPackets     = 0
)

func listenClients() {
	buffer := make([]byte, 1024)
	address, err := net.ResolveUDPAddr("udp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
	}
	server, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
	for {
		n, clientAddr, _ := server.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		receivedPackets++
		packetsLog("pacote recebido de: ", clientAddr, receivedPackets, sentPackets, n)
		fmt.Printf("%s\n", buffer)
		go sendToClient(server, clientAddr)
	}
}

func sendToClient(connection *net.UDPConn, clientAddr *net.UDPAddr) {
	n, err := connection.WriteToUDP([]byte("E ai cliente\n"), clientAddr)
	if err != nil {
		fmt.Println(err)
	}
	sentPackets++
	packetsLog("pacote enviado para: ", clientAddr, receivedPackets, sentPackets, n)
}

func packetsLog(msg string, addr *net.UDPAddr, rcvdTotal, sndTotal, n int) {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rcvd := strconv.Itoa(rcvdTotal)
	snd := strconv.Itoa(sndTotal)
	log.SetOutput(f)
	log.Println(msg, addr, "/total enviado:", snd, "/total recebido:", rcvd, "/tamanho da mensagem:", n)
}
