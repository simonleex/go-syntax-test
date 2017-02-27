package main

import (
	"net/http"
	"fmt"
	"runtime"
)


func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello,World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
