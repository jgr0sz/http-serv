package main

import (
	"log"
	"net/url"
)

func main() {
	routes = nil
	addRoute("GET", "/hello", func(req *Request) *Response {
		return NewResponse(200, "OK", nil, "HELLO!")
	})

	addRoute("GET", "/login", func(req *Request) *Response {
    html := `<!DOCTYPE html>
<html>
    <head><title>Login</title></head>
    <body>
        <form method="POST" action="/login">
            <input type="text" name="username" placeholder="Username"/>
            <input type="password" name="password" placeholder="Password"/>
            <button type="submit">Login</button>
        </form>
    </body>
</html>`
    return NewResponse(200, "OK", map[string]string{
        "Content-Type": "text/html"}, html)
	})

	addRoute("POST", "/login", func(req *Request) *Response {
		params, err := url.ParseQuery(req.body)
		if err != nil {
			return NewResponse(400, "Bad Request", nil, "Invalid form data.")
		}

		username := params.Get("username")
		password := params.Get("password")

		if username == "" || password == "" {
			return NewResponse(400, "Bad Request", nil, "Username and password are required.")
		}

		//db lookup logic...
		log.Printf("Login attempt: %s", username)

		return NewResponse(200, "OK", map[string]string{
			"Content-Type": "text/html",
		}, "<h1>Welcome, " + username + "!</h1>")
	})
	listenTCP("127.0.0.1:9999")

}
