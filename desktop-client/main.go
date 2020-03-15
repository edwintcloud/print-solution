package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		parseResult := checkMessage(message)

		// send parseResult back to client
		conn.Write([]byte(parseResult + "\n"))
	}
}

func checkMessage(message string) string {
	messageParts := strings.Split(message, " ")
	if len(messageParts) > 2 {

		switch messageParts[0] {
		case "PRNT":
			log.Printf("received PRNT opcode: sending to print service\n")
			return handlePrint(messageParts[1], strings.Join(messageParts[2:], " "))
		default:
			log.Printf("received %s opcode: not implemented", messageParts[0])
			return message
		}
	}

	return message
}

func handlePrint(printer, data string) string {
	c1 := exec.Command("echo", data)
	c2 := exec.Command("lp")
	if printer != "DEFAULT" {
		c2 = exec.Command("lp", "-d", printer)
	}

	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	c2.Start()
	c1.Run()
	c2.Wait()

	return fmt.Sprintf("successfully printed data to %s", printer)
}
