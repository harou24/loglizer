package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/analysis", analysisHandler)
	log.Println("Server starting on port 15442...")
	log.Fatal(http.ListenAndServe(":15442", nil))
}

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	_, err = io.Copy(ioutil.Discard, file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte("data received\n"))
	if err != nil {
		log.Printf("failed to write to response: %s", err)
		return
	}
}
