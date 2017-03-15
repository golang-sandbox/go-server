package main

import (
	"io"
	"net/http"
	"log"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
  const port = ":8080"
	// http.HandleFunc("/hello", HelloServer)
	// log.Fatal(http.ListenAndServe(port, nil))
  log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir("/Volumes/Martial-Soul/coding2/go-workspace/test-server/src"))))
  // fmt.Printf("Listening on port: ", port)
}
