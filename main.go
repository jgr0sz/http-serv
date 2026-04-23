package main

// "macro" consts for common status codes
const (
	OK          = "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND   = "HTTP/1.1 404 Not Found\r\n\r\n"
	NOT_ALLOWED = "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
)

func main() {
	listenTCP("127.0.0.1:9999")
}
