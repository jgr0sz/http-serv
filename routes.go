package main

//Struct representation of an HTTP response
type Response struct {
	status string
	headers map[string]string
	body string
}