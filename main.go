package main

import (
	"net/http"
)

func main() {
	panic(http.ListenAndServe("127.0.0.1:1234", http.FileServer(http.Dir("./"))))
}
