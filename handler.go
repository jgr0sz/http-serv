package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

//Struct representation of an HTTP request
type Request struct {
	method string
	path string
	proto string
	headers map[string]string
	body string
}

//Takes the received request and parses it into struct Request
func parseRequest(reader *bufio.Reader) (*Request, error) {
	reqLine, err := reader.ReadString('\n');
	if err != nil {
		return nil, fmt.Errorf("Unable to read start line of request: %w", err)
	}
	
	//Splits reqLine into method/path/proto (uses whitespace as delimiter)
	reqLineParts := strings.Fields(reqLine)
	//Invalid start line check (first line of the request)
	if len(reqLineParts) < 3 {
		return nil, fmt.Errorf("Malformed request start line: %s", reqLine)
	}

	request := &Request{
		method: reqLineParts[0],
		path: reqLineParts[1], 
		proto: reqLineParts[2],
		headers: map[string]string{},
	}

	//Parsing the remaining headers of the request
	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Unable to read header: %w", err)
		}

		//End of header
		if headerLine == "\r\n" {
			break
		}

		//Splitting headers using SplitN() to obtain two substrings for our key/value map
		headerParts := strings.SplitN(headerLine, ": ", 2)
		if len(headerParts) != 2 {
			return nil, fmt.Errorf("Malformed header: %v", headerParts)
		}
		//Trimming spaces, we add our header parts to our request's headers map
		request.headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
	}

	//parsing the request body (if present)
	if length, exists := request.headers["Content-Length"]; exists {
		contentLength, err := strconv.Atoi(length)
		if err != nil {
			return nil, fmt.Errorf("Invalid Content-Length header: %w", err)
		}

		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)

		if err != nil {
			return nil, fmt.Errorf("Unable to read body: %w", err)
		}
		request.body = string(body)
	}
	return request, nil
}

//Parses HTTP request into a Request object and responds per user-defined routes.
func connHandler(conn net.Conn) {
	defer conn.Close()
	//Connection timeout
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	//Parses request recieved
	reader := bufio.NewReader(conn)
	request, err := parseRequest(reader)
	if err != nil {
		log.Printf("Bad request from %s: %v", conn.RemoteAddr(), err)
	}
	fmt.Print(request)
}
