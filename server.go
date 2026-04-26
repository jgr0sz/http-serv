package main

import (
	"log"
	"net"
)

// Main listener function for TCP connections to the server
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

	//Blocks on Accept() until a connection is found. Connection queue is handled concurrently
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection ", err)
			continue
		}
		go connHandler(conn)
	}
}
