package main

import (
	"fmt"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is UP and Running!"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	fmt.Println("url-shortner-service application started...")
	http.HandleFunc("/health", handleHealth)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
