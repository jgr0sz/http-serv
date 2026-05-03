# http-serv
A simple, bespoke web server in Go

#### Currently it functions as follows:

- server contains a route registry that maps handler functions to methods/paths
- listenTCP() receives an address to start listening to TCP connections on
- blocks listenTCP() loop on Accept() until a connection is found
- connections are sent to connHandler() asynchronously via goroutine
- connHandler() receives the HTTP request and parses it using parseRequest()
- parsed data is sent through the route registry to invoke matching handlers or return an error
