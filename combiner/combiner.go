package combiner

import (
	"log"
	"os"
)

func WriteProcessedLogsToFile(fileName string, processedLogsChan <-chan string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer file.Close()

	for logEntry := range processedLogsChan {
		if _, err := file.WriteString(logEntry + "\n"); err != nil {
			log.Fatalf("failed to write to file: %s", err)
		}
	}
}
