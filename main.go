package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := "8080"
	dir, err := filepath.Abs("./")
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	if err == nil {
		fmt.Printf("Serving HTTP on 0.0.0.0 port %s ...\n", port)
		err = http.ListenAndServe(":"+port, http.FileServer(http.Dir(dir)))
	}
	panic(err)
}
