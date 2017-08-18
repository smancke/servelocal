package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	listen := "127.0.0.1:1234"
	if len(os.Args) > 1 {
		listen = os.Args[1]
	}

	fileServer := http.FileServer(http.Dir("./"))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		fileServer.ServeHTTP(w, r)
	})

	go openBrowser("http://" + listen)

	fmt.Printf("start serving http://%v\n", listen)
	panic(http.ListenAndServe(listen, handler))
}

func openBrowser(url string) {
	err := exec.Command("x-www-browser", url).Run()
	if err != nil {
		fmt.Printf("error starting browser: %v\n", err)
	}
}
