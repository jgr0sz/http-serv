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

//Struct encapsulating the contents of HTTP requests for use.
type Request struct {
	//Method, path, and proto are all in the first line
	method string
	path string
	proto string
	//key-value pairs for the headers since they are naturally that way (e.g. Content-Length: 64)
	headers map[string]string
	//stores the contents of Content-Length, empty otherwise
	body string
}

//Takes the received request and parses it into struct Request
func parseRequest(reader *bufio.Reader) (*Request, error) {
	//Read first line for method/path/proto
	reqLine, err := reader.ReadString('\n');
	if err != nil {
		log.Print("Unable to read start line of request: %w", err)
	}
	
	//Splits reqLine into method/path/proto using their whitespace as a delimiter
	reqLineParts := strings.Fields(reqLine)
	//Invalid start line check (first line of the request)
	if len(reqLineParts) < 3 {
		return nil, fmt.Errorf("Malformed request start line: %s", reqLine)
	}

	//Initializing and assigning the already parsed parts of our request
	request := &Request{
		method: reqLineParts[0],
		path: reqLineParts[1], 
		proto: reqLineParts[2],
		headers: map[string]string{},
		
	}

	//while-loop to parse the remaining headers of the request
	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Unable to read header: %w", err)
		}
		log.Printf("header line!!!!: %q", headerLine)

		//End of header
		if headerLine == "\r\n" {
			break
		}

		//Using : as a delimiter, we use SplitN() to obtain two substrings for our key/value map
		headerParts := strings.SplitN(headerLine, ": ", 2)
		if len(headerParts) != 2 {
			return nil, fmt.Errorf("Malformed header: %v", headerParts)
		}
		//Trimming spaces, we add our header parts to our request's headers map
		request.headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
	}

	//parsing the request body (if present); 
	if length, exists := request.headers["Content-Length"]; exists {
		contentLength, err := strconv.Atoi(length)
		if err != nil {
			return nil, fmt.Errorf("Invalid Content-Length header: %w", err)
		}

		//Dynamic byte array for our body (JSON, XML, etc. etc.) we fill with what would be the body
		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		request.body = string(body)
	}
	return request, nil
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	//Connection timeout (5s)
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	reader := bufio.NewReader(conn)
	parseRequest(reader)
}
