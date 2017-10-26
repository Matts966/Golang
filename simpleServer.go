package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("your request is as follows\n"))
	w.Write([]byte(r.Method))
	w.Write([]byte(r.Host))
	w.Write([]byte(r.RequestURI))
	w.Write([]byte(r.RemoteAddr))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}