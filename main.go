package main

import (
	"log"
	"net/http"
)

func main() {
	var port = "5050"
	http.HandleFunc("/", home)
	log.Println("Server is ready to handle requests at port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("User-Agent:", r.UserAgent())
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	// response := "ok"
	response := http.StatusOK
	_, err := w.Write([]byte(response))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
