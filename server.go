package main

import (
	"log"
	"net"
)

// Main listener function for TCP connections to the server
func listenTCP() {
	//TCP listener creation
	listener, err := net.Listen("tcp", "127.0.0.1:9999")

	//Error check
	if err != nil {
		log.Fatal("Error while listening ", err)
	}

	//Logging listener status
	log.Printf("\nListening on\n------------\nType: %s\nAddress: %s\n",
	listener.Addr().Network(),
	listener.Addr().String())
	//deferred statement to close our listener at the end of the function
	defer listener.Close()

	//Loop that blocks on Accept() until a connection is found. After error checking, the connection
	//is handled concurrently through a goroutine.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection ", err)
			continue
		}
		go connHandler(conn)
	}
}