package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// "macro" consts for common status codes
const (
	OK = "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND = "HTTP/1.1 404 Not Found\r\n\r\n"
	NOT_ALLOWED = "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
)

// Defines a serverside response to an HTTP request after performing error checking
func writeResponse(conn net.Conn, response string) {
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

// Extracts the first line of the HTTP request to determine its type/validity
func connHandler(conn net.Conn) {
	defer conn.Close()

	//Instantiation of a reader to go through conn
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading request: ", err)
		return
	}

	//Splits header into sections which we will utilize
	headerSections := strings.Fields(line)
	if len(headerSections) < 3 {
		log.Println("Malformed request: ", line)
		return
	}

	//GET request case (rudimentary)
	if headerSections[0] == "GET" {
		writeResponse(conn, OK + "GET Response of some kind....")
	} else {
		writeResponse(conn, NOT_ALLOWED + "Response type not allowed.")
	}
}

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

func main() {
	listenTCP()
}
