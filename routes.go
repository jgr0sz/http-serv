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

//Helper function to construct responses 
func NewResponse(status int, statusText string, headers map[string]string, body string) *Response {
	//Default headers
	if headers == nil {
		headers = map[string]string {"Content-Type": "text/plain"}
	}
	return &Response{
		status: status,
		statusText: statusText,
		headers: headers,
		body: body,
	}
}

//Adds routes to registry
func addRoute(method, path string, handler HandlerFunc) {
	routes = append(routes, Route{method, path, handler})
}

//Checks request data, providing a response by invoking the stored route
func invokeRoute(req *Request) *Response {
	pathMatch := false
	for _, route := range routes {
		if route.path == req.path {
			pathMatch = true
			if route.method == req.method {
				return route.handler(req)
			}
		}
	}
	if pathMatch {
		 NewResponse(405, "Method Not Allowed", nil, "Method not allowed.")
	}
	return NewResponse(404, "Not Found", nil, "The resource you were looking for was not found.")
}

