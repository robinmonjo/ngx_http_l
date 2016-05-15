package main

import (
	"fmt"
	"net/http"
)

func startApi(port string) error {
	http.HandleFunc("/", handle)
	return http.ListenAndServe(":"+port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here you will access greatness my friend\n")
}
