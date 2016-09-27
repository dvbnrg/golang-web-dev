package main

import (
	"io"
	"net/http"
)

type myHandler int
func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	}
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
