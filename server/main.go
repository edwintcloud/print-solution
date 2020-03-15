package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	serverAddr = "127.0.0.1:8081"
)

func main() {
	conn, _ := net.Dial("tcp", serverAddr)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}
