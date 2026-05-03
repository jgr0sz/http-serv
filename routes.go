package main

type Response struct {
	status     int
	statusText string
	headers    map[string]string
	body       string
}

type HandlerFunc func(*Request) *Response

type Route struct {
	method  string
	path    string
	handler HandlerFunc
}

var routes []Route

//Adds routes to registry
func addRoute(method, path string, handler HandlerFunc) {
	routes = append(routes, Route{method, path, handler})
}

//Checks request data, providing a response by invoking the stored route
func invokeRoute(req *Request) *Response {
	return &Response{}
}
