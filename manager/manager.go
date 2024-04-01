package manager

import (
	"bufio"
	"loglizer/processor"
	"loglizer/reader"
	"runtime"
	"sync"
)

const (
	logsChanSize          = 10000000
	processedLogsChanSize = 1000000
)

func StartLogProcessingWorkflow(scanner *bufio.Scanner) <-chan string {
	logEntriesChan := make(chan []string, logsChanSize)
	processedPairsChan := make(chan string, processedLogsChanSize)

	var wg sync.WaitGroup

	workerCount := runtime.NumCPU()
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go processor.SummarizeLogFrequency(logEntriesChan, processedPairsChan, &wg)
	}

	go func() {
		reader.ReadHourlyLogBatches(scanner, logEntriesChan)
		close(logEntriesChan)
	}()

	go func() {
		defer close(processedPairsChan)
		wg.Wait()
	}()

	return processedPairsChan
}
