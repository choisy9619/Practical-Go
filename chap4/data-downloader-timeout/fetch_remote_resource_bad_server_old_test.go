package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func startBadTestHTTPServerV2(shutdownServer chan struct{}) *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				<-shutdownServer
				fmt.Fprint(w, "Hello World")
			}))
	return ts
}
