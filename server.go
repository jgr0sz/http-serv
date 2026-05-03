package main

import (
	"fmt"
	"log"
	"net"
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
func writeResponse(conn net.Conn, status int, statusText string, body string) {
	response := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s",
		status, statusText, len(body), body,
	)
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Printf("Unable to write HTTP response, contents:\n%s", response)
	}
}
