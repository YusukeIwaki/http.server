package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type LoggingHandler struct {
	Handler http.Handler
}

func (loggingHandler *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s\t%s\t\t\"%s %s\"\n", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Method, r.URL.Path)
    loggingHandler.Handler.ServeHTTP(w, r)
}

func main() {
	port := "8080"
	dir, err := filepath.Abs("./")
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	if err == nil {
		fmt.Printf("Serving HTTP on 0.0.0.0 port %s ...\n", port)
		handler := LoggingHandler { http.FileServer(http.Dir(dir)) }
		err = http.ListenAndServe(":"+port, &handler)
	}
	panic(err)
}
