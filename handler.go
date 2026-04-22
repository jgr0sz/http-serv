package main

import (
	"bufio"
	"log"
	"net"
	"strings"
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