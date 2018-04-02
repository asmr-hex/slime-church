package main

import (
	"fmt"
	"net/http"
)

const (
	IP = "192.168.50.145"
)

func main() {
	http.HandleFunc("/", SayHi)
	fmt.Println("listening on slime world")
	http.ListenAndServe(IP+":6666", nil)
}

func SayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hii"))
}
