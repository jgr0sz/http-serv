package main

//Struct representation of an HTTP response
type Response struct {
	status     int
	statusText string
	headers    map[string]string
	body       string
}

//Function that defines the behavior at a certain endpoint
type HandlerFunc func(*Request) *Response

//Struct representation of a route
type Route struct {
	method  string
	path    string
	handler HandlerFunc
}
