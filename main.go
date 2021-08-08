package main

import (
	"fmt"
	"net/http"

	"example.com/webcache/webcache"
)

func main() {
	http.HandleFunc("/", webcache.CachedWebpageHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Server failed to start", err)
	}
}
