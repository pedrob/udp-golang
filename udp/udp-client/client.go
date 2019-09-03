package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	buffer := make([]byte, 1024)
	connection, _ := net.Dial("udp", "localhost:8888")
	defer connection.Close()
	fmt.Fprintf(connection, "E ai servidor")
	bufio.NewReader(connection).Read(buffer)
	fmt.Printf("%s\n", buffer)
}
