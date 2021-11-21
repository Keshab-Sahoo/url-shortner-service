package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortner-service/src"
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

func handleShortUrl(writer http.ResponseWriter, req *http.Request) {
	if "GET" != req.Method {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		originalURL := req.URL.Query().Get("longURL")
		host := req.Host
		shortURL := src.GenerateShortURL()
		src.SaveToFile(shortURL, originalURL)
		resp := src.BuildURLResponse(host, shortURL, originalURL)
		jsonBytes, err := json.Marshal(resp)
		if err != nil {
			writer.Write([]byte("Failed to generate response"))
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(jsonBytes)
	}
}

func handleRedirect(writer http.ResponseWriter, req *http.Request) {
	//remove first character '/' from uri
	shortPath := req.URL.Path[1:]
	newUrl, err := src.GetOriginalURL(shortPath)
	if err != nil {
		panic(err)
	}
	http.Redirect(writer, req, newUrl, http.StatusSeeOther)
}

func main() {
	fmt.Println("url-shortner-service application started...")
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/short-url", handleShortUrl)
	http.HandleFunc("/", handleRedirect)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
