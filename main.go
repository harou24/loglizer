package main

import (
	"bufio"
	"log"
	"loglizer/processor"
	"loglizer/reader"
	"net/http"
	"runtime"
	"sync"
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

	logEntriesChan := make(chan []string, 10000000)
	processedPairsChan := make(chan string, 100)

	var wg sync.WaitGroup

	workerCount := runtime.NumCPU()
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go processor.SummarizeLogFrequency(logEntriesChan, processedPairsChan, &wg)
	}

	scanner := bufio.NewScanner(file)
	go func() {
		reader.ReadHourlyLogBatches(scanner, logEntriesChan)
		close(logEntriesChan)
	}()

	go func() {
		defer close(processedPairsChan)
		wg.Wait()
	}()

	w.Header().Set("Content-Type", "text/csv")
	for pair := range processedPairsChan {
		if _, err := w.Write([]byte(pair + "\n")); err != nil {
			log.Printf("failed to write to response: %s", err)
			return
		}
	}
}
