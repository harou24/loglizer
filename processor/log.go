package processor

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	File      string
	Message   string
}

func SummarizeLogFrequency(logEntriesChan <-chan []string, processedPairsChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for lines := range logEntriesChan {
		var entries []LogEntry
		for _, line := range lines {
			entry, err := parseLogEntry(line)
			if err != nil {
				log.Printf("error parsing log entry: %s", err)
				continue
			}
			entries = append(entries, entry)
		}
		dateHour, mostFrequent := findMostFrequentPairAndFormat(entries)
		processedPairsChan <- fmt.Sprintf("%s,%s", dateHour, mostFrequent)
	}
}

func findMostFrequentPairAndFormat(entries []LogEntry) (string, string) {
	frequencyMap := make(map[string]int)
	var maxCount int
	var mostFrequentPair string
	var timestamp time.Time

	for _, entry := range entries {
		pair := fmt.Sprintf("%s,%s", entry.File, entry.Message)
		frequencyMap[pair]++
		if frequencyMap[pair] > maxCount {
			maxCount = frequencyMap[pair]
			mostFrequentPair = pair

			// Keep track of the timestamp for formatting
			timestamp = entry.Timestamp
		}
	}

	// New date and hour format: MMDDYYYY,HH
	dateHour := timestamp.Format("01022006,15")
	return dateHour, mostFrequentPair
}

func parseLogEntry(line string) (LogEntry, error) {
	parts := strings.SplitN(line, ",", 3)
	if len(parts) != 3 {
		return LogEntry{}, fmt.Errorf("invalid log entry: %s", line)
	}
	timestamp, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return LogEntry{}, err
	}
	return LogEntry{
		Timestamp: timestamp,
		File:      parts[1],
		Message:   parts[2],
	}, nil
}
