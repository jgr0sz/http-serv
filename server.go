package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

//Main listener function for TCP connections to the server
func listenTCP(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error while listening ", err)
	}

	//Logging listener status
	log.Printf("\nListening on\n------------\nType: %s\nAddress: %s\n",
		listener.Addr().Network(),
		listener.Addr().String())
	defer listener.Close()

	//Blocks until a connection is found, conn queue is handled concurrently
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection ", err)
			continue
		}
		go connHandler(conn)
	}
}

//Helper function for HTTP responses
func writeResponse(conn net.Conn, response *Response) {
	//Constructing response string
	var sb strings.Builder

	fmt.Fprintf(&sb, "HTTP/1.1 %d %s\r\n", response.status, response.statusText)
	for k, v := range response.headers {
		fmt.Fprintf(&sb, "%s: %s\r\n", k, v)
	}
	fmt.Fprintf(&sb, "Content-Length: %d\r\n", len(response.body))
	fmt.Fprintf(&sb, "\r\n%s", response.body)

	_, err := conn.Write([]byte(sb.String()))
	if err != nil {
		log.Printf("Unable to write HTTP response, contents:\n%s", sb.String())
	}
}
