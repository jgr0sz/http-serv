package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseRequest_GET(t *testing.T) {
    raw := "GET /index.html HTTP/1.1\r\n" +
           "Host: example.com\r\n" +
           "\r\n"

    reader := bufio.NewReader(strings.NewReader(raw))
    req, err := parseRequest(reader)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if req.method != "GET" {
        t.Errorf("expected method GET, got %s", req.method)
    }
    if req.path != "/index.html" {
        t.Errorf("expected path /index.html, got %s", req.path)
    }
    if req.headers["Host"] != "example.com" {
        t.Errorf("expected Host example.com, got %s", req.headers["Host"])
    }
}

func TestParseRequest_POST(t *testing.T) {
    raw := "POST /login HTTP/1.1\r\n" +
           "Host: example.com\r\n" +
           "Content-Type: application/x-www-form-urlencoded\r\n" +
           "Content-Length: 27\r\n" +
           "\r\n" +
           "username=john&password=1234"

    reader := bufio.NewReader(strings.NewReader(raw))
    req, err := parseRequest(reader)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if req.method != "POST" {
        t.Errorf("expected method POST, got %s", req.method)
    }
    if req.path != "/login" {
        t.Errorf("expected path /login, got %s", req.path)
    }
    if req.body != "username=john&password=1234" {
        t.Errorf("expected body username=john&password=1234, got %s", req.body)
    }
}

func TestParseRequest_POST_MismatchedContentLength(t *testing.T) {
    raw := "POST /login HTTP/1.1\r\n" +
           "Host: example.com\r\n" +
           "Content-Length: 999\r\n" + //wrongly-sized Content-Length
           "\r\n" +
           "username=john&password=1234"

    reader := bufio.NewReader(strings.NewReader(raw))
    _, err := parseRequest(reader)

    if err == nil {
        t.Fatal("expected error for mismatched Content-Length, got nil")
    }
}

func TestParseRequest_POST_InvalidContentLength(t *testing.T) {
    raw := "POST /login HTTP/1.1\r\n" +
           "Host: example.com\r\n" +
           "Content-Length: abc\r\n" +
           "\r\n" +
           "username=john&password=1234"

    reader := bufio.NewReader(strings.NewReader(raw))
    _, err := parseRequest(reader)

    if err == nil {
        t.Fatal("expected error for invalid Content-Length, got nil")
    }
}