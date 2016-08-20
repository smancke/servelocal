package main

import (
	"net/http"
	"os"
)

func main() {
	listen := "127.0.0.1:1234"
	if len(os.Args) > 1 {
		listen = os.Args[1]
	}

	handler := http.FileServer(http.Dir("./"))

	panic(http.ListenAndServe(listen, handler))
}
