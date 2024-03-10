package main

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := `Usage: mync [http|grpc] -h
http: A HTTP client.
http : <options> server

Options:
	-verb string
	HTTP method (defualt "GET")

grpc: A gRPC client.
grpc : <options> server

Options:
	-body string
	Body of request
	-method string
	Method to call
`
	byteBuf := new(bytes.Buffer)
}
