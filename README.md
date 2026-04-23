# http-serv
A simple, bespoke web server in Go

#### Currently it functions as follows:

- listenTCP() receives an address to start listening to TCP connections on
- blocks loop inside listenTCP() until a connection is found
- connection is handled through an asynchronous call to connHandler()
- connHandler() receives the HTTP request through the connection and parses it using parseRequest()
- parsed data is routed to individual handlers based on method type (getHandler, postHandler, etc.)
- handlers check endpoints; depending on the method and endpoint, data is sent back in the form of a response
