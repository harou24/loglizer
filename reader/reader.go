package reader

import (
	"bufio"
	"log"
	"strings"
	"time"
)

func ReadHourlyLogBatches(scanner *bufio.Scanner, logEntriesChan chan<- []string) {
	var lines []string
	lastHour := -1
	for scanner.Scan() {
		line := scanner.Text()

		timestamp, err := time.Parse(time.RFC3339, strings.SplitN(line, ",", 2)[0])
		if err != nil {
			log.Printf("error parsing timestamp: %s", err)
			continue
		}

		entryHour := timestamp.Hour()
		if lastHour != -1 && entryHour != lastHour {
			logEntriesChan <- lines
			lines = nil
		}
		lines = append(lines, line)
		lastHour = entryHour
	}
	if len(lines) > 0 {
		logEntriesChan <- lines
	}
}
