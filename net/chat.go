package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incommingClients = make(chan Client)
	leavingClients   = make(chan Client)
	messages         = make(chan string)
)

var (
	port = flag.Int("p", 3090, "puerto")
	host = flag.String("h", "localhost", "host")
)

// assign connection to a server
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	// create a channel to send messages
	message := make(chan string)
	// function to send messages troghput the connection
	go MessageWrite(conn, message)

	// set client name and receive welcome message
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server, your name is %s \n" + clientName)
	messages <- fmt.Sprintf("New Client is here, name %s \n", clientName)
	// set the channel for new clients
	incommingClients <- message

	inputMessage := bufio.NewScanner(conn)
	// iterate while the terminal is open in order to check for new messages.
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s : %s \n", clientName, inputMessage.Text())
	}
	// set the channel for clients that are leaving and send a goodbye message
	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodbye!\n", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for messaage := range messages {
		fmt.Println(conn, messaage)
	}
}
