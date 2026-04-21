package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// "macro" consts for common status codes
const (
	OK_GET      = "HTTP/1.1 200 OK\r\n\r\n"
	BAD_GET     = "HTTP/1.1 404 Not Found\r\n\r\n"
	NOT_ALLOWED = "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
)

// Extracts the first line of the HTTP request to determine its type/validity
func connHandler(conn net.Conn) {
	defer conn.Close()

	//Instantiation of a reader to go through conn
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error reading request: ", err)
	}

	//Splits header into sections which we will utilize
	headerSections := strings.Fields(line)
	if len(headerSections) < 3 {
		fmt.Print("Malformed request: ", line)
	}

	if headerSections[0] == "GET" {
		conn.Write([]byte(strings.Join([]string{OK_GET, "Here is a GET request"}, " ")))
	} else {
		conn.Write([]byte(strings.Join([]string{NOT_ALLOWED, "Request type not allowed..."}, " ")))
	}
}

// Main listener function for TCP connections to the server
func listenTCP() {
	//Listening for a TCP response
	response, err := net.Listen("tcp", "127.0.0.1:9999")

	//If a response was found, output it, output any error otherwise
	if err != nil {
		fmt.Print("Error while listening ", err)
		return
	}
	//deferred statement to close our listener at the end of the function
	defer response.Close()

	//Message to indicate connection
	fmt.Printf("Connected\n-------\nType: %s\nAddress: %s",
		response.Addr().Network(),
		response.Addr().String())

	conn, err := response.Accept()
	if err != nil {
		fmt.Print("Unable to accept connection ", err)
	}
	connHandler(conn)
}

func main() {
	//Gradually increase wait time after listens
	for {
		listenTCP()
	}
}
