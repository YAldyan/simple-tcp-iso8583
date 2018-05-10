package main

import (
	"mofax/iso8583"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	specFile string = "mofax/iso8583/spec21987.yml"
)

func main() {
	tcpServer()
}

func tcpServer() {

	fmt.Println("Launching server...")

	service := ":8081"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

	checkError(err)

	// listen on all interfaces
	ln, _ := net.ListenTCP("tcp", tcpAddr)

	checkError(err)

	// accept connection on port
	// conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {

		conn, err := ln.Accept()

		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		fmt.Print("Message Received:", string(message))

		// sample process for string received
		newmessage := strings.ToUpper(message)

		if err != nil {
			continue
		}

		// concurrent access
		go handleClient(conn, newmessage)

	}
}

func handleClient(conn net.Conn, msg string) {
	defer conn.Close()
	conn.Write([]byte(msg))
	// don't care about return value
	// we're finished with this client
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func displayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func unpacked(in string) {

	isostruct := iso8583.NewISOStruct(specFile, true)
	parsed, err := isostruct.Parse(in)

	displayError(err)

	fmt.Printf("%#v, %#v\n%#v", parsed.Mti, parsed.Bitmap, parsed.Elements)
}
