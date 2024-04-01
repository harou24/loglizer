package combiner

import (
	"log"
	"os"
)

func WriteProcessedLogsToFile(fileName string, processedPairsChan <-chan string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer file.Close()

	for pair := range processedPairsChan {
		if _, err := file.WriteString(pair + "\n"); err != nil {
			log.Fatalf("failed to write to file: %s", err)
		}
	}
}
