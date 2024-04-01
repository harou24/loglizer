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
	processedLogsChan := make(chan string, processedLogsChanSize)

	var wg sync.WaitGroup

	workerCount := runtime.NumCPU()
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go processor.SummarizeLogFrequency(logEntriesChan, processedLogsChan, &wg)
	}

	go func() {
		reader.ReadHourlyLogBatches(scanner, logEntriesChan)
		close(logEntriesChan)
	}()

	go func() {
		defer close(processedLogsChan)
		wg.Wait()
	}()

	return processedLogsChan
}
