package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type MyServer struct{}

func (ms *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from my server")
}

type LoggedServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (ls *LoggedServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(ls.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(ls.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(ls.LogWriter, "Content lenght: %d\n", r.ContentLength)
	fmt.Fprintf(ls.LogWriter, "Method: %s\n", r.Method)
	ls.Handler.ServeHTTP(w, r)
}

func main() {
	http.Handle("/", &LoggedServer{
		Handler:   &MyServer{},
		LogWriter: os.Stdout,
	})
	http.ListenAndServe(":8080", nil)
}
