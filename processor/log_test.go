package processor

import (
	"sync"
	"testing"
	"time"
)

func TestParseLogEntry(t *testing.T) {
	input := "2019-04-30T12:01:39+02:00,network.go,Network connection established"
	expectedTimestamp, _ := time.Parse(time.RFC3339, "2019-04-30T12:01:39+02:00")
	expected := LogEntry{
		Timestamp: expectedTimestamp,
		File:      "network.go",
		Message:   "Network connection established",
	}
	result, err := parseLogEntry(input)
	if err != nil {
		t.Errorf("parseLogEntry returned an error: %v", err)
	}

	if !result.Timestamp.Equal(expected.Timestamp) || result.File != expected.File || result.Message != expected.Message {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestFindMostFrequentPairAndFormat(t *testing.T) {
	// The most frequent message is "Error: Meme generator ran out of memes" in the test entries batch
	expectedPair := "memeGenerator.go,Error: Meme generator ran out of memes"

	expectedDateHour := "04302019,12"

	dateHour, mostFrequentPair := findMostFrequentPairAndFormat(logEntries)

	if dateHour != expectedDateHour || mostFrequentPair != expectedPair {
		t.Errorf("findMostFrequentPairAndFormat() = %v, %v; want %v, %v", dateHour, mostFrequentPair, expectedDateHour, expectedPair)
	}
}

func TestSummarizeLogFrequency(t *testing.T) {
	// Setup: Create channels and a WaitGroup
	logEntriesChan := make(chan []string)
	processedPairsChan := make(chan string, 10)
	var wg sync.WaitGroup

	mockLines := []string{
		"2019-04-30T12:01:39+02:00,network.go,Network connection established",
		"2019-04-30T12:01:42+02:00,db.go,Transaction failed",
		"2019-04-30T12:06:19+02:00,memeGenerator.go,Error: Meme generator ran out of memes",
		"2019-04-30T12:06:19+02:00,memeGenerator.go,Error: Meme generator ran out of memes",
	}

	// Expected output
	expectedOutput := "04302019,12,memeGenerator.go,Error: Meme generator ran out of memes"

	// Start the function in a goroutine
	wg.Add(1)
	go SummarizeLogFrequency(logEntriesChan, processedPairsChan, &wg)

	// Send mock data to the channel
	go func() {
		logEntriesChan <- mockLines
		close(logEntriesChan)
	}()

	// Wait for the processing to complete
	wg.Wait()
	close(processedPairsChan)

	// Check the output
	receivedOutput, ok := <-processedPairsChan
	if !ok {
		t.Fatal("Expected an output from processedPairsChan, but it was closed without sending any data")
	}

	if receivedOutput != expectedOutput {
		t.Errorf("SummarizeLogFrequency output = %v, want %v", receivedOutput, expectedOutput)
	}
}

var logEntries = []LogEntry{
	{
		Timestamp: time.Date(2019, 4, 30, 12, 1, 39, 0, time.FixedZone("CET", 2*60*60)),
		File:      "network.go",
		Message:   "Network connection established",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 1, 42, 0, time.FixedZone("CET", 2*60*60)),
		File:      "db.go",
		Message:   "Transaction failed",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 2, 3, 0, time.FixedZone("CET", 2*60*60)),
		File:      "tardis.go",
		Message:   "TARDIS dematerializing",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 2, 44, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 5, 26, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 5, 56, 0, time.FixedZone("CET", 2*60*60)),
		File:      "procrastinationService.go",
		Message:   "Error: Failed to procrastinate",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 6, 19, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 7, 20, 0, time.FixedZone("CET", 2*60*60)),
		File:      "hal9000.go",
		Message:   "Error: I know that you and Frank were planning to disconnect me",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 8, 3, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 8, 8, 0, time.FixedZone("CET", 2*60*60)),
		File:      "context.go",
		Message:   "Error: Failed to create context",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 8, 23, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 8, 38, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 10, 43, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 11, 41, 0, time.FixedZone("CET", 2*60*60)),
		File:      "cache.go",
		Message:   "Cache created",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 11, 41, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 11, 53, 0, time.FixedZone("CET", 2*60*60)),
		File:      "tardis.go",
		Message:   "TARDIS dematerializing",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 15, 2, 0, time.FixedZone("CET", 2*60*60)),
		File:      "rubberDuckDebugger.go",
		Message:   "Rubber duck debugging started",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 15, 3, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 16, 52, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 16, 53, 0, time.FixedZone("CET", 2*60*60)),
		File:      "jokeGenerator.go",
		Message:   "Error: Failed to generate a joke",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 16, 53, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 17, 52, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 18, 4, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 18, 6, 0, time.FixedZone("CET", 2*60*60)),
		File:      "coffeeMachine.go",
		Message:   "Coffee is ready",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 18, 17, 0, time.FixedZone("CET", 2*60*60)),
		File:      "hal9000.go",
		Message:   "Error: I know that you and Frank were planning to disconnect me",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 18, 24, 0, time.FixedZone("CET", 2*60*60)),
		File:      "rubberDuckDebugger.go",
		Message:   "Rubber duck debugging started",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 19, 47, 0, time.FixedZone("CET", 2*60*60)),
		File:      "starTrek.go",
		Message:   "Error: Failed to beam up",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 19, 56, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 20, 20, 0, time.FixedZone("CET", 2*60*60)),
		File:      "cache.go",
		Message:   "Cache created",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 20, 25, 0, time.FixedZone("CET", 2*60*60)),
		File:      "network.go",
		Message:   "Network connection established",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 21, 37, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 21, 39, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 21, 43, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 21, 45, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 22, 47, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 23, 59, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 26, 13, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 28, 28, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 29, 3, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 29, 33, 0, time.FixedZone("CET", 2*60*60)),
		File:      "cache.go",
		Message:   "Cache created",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 32, 15, 0, time.FixedZone("CET", 2*60*60)),
		File:      "tardis.go",
		Message:   "TARDIS dematerializing",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 32, 17, 0, time.FixedZone("CET", 2*60*60)),
		File:      "hal9000.go",
		Message:   "Error: I know that you and Frank were planning to disconnect me",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 34, 5, 0, time.FixedZone("CET", 2*60*60)),
		File:      "db.go",
		Message:   "Transaction failed",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 34, 46, 0, time.FixedZone("CET", 2*60*60)),
		File:      "cache.go",
		Message:   "Cache created",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 34, 46, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 34, 57, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 35, 7, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 36, 12, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 37, 46, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 38, 58, 0, time.FixedZone("CET", 2*60*60)),
		File:      "db.go",
		Message:   "Transaction failed",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 39, 46, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 40, 12, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 40, 17, 0, time.FixedZone("CET", 2*60*60)),
		File:      "network.go",
		Message:   "Network connection established",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 41, 3, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 42, 6, 0, time.FixedZone("CET", 2*60*60)),
		File:      "hal9000.go",
		Message:   "Error: I know that you and Frank were planning to disconnect me",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 42, 29, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 44, 7, 0, time.FixedZone("CET", 2*60*60)),
		File:      "client.go",
		Message:   "Error: Failed to connect to client",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 45, 25, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authorization.go",
		Message:   "Error: User is not authorized",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 45, 53, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 45, 56, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 46, 48, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 47, 35, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 47, 49, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 48, 1, 0, time.FixedZone("CET", 2*60*60)),
		File:      "network.go",
		Message:   "Network connection established",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 48, 2, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 48, 36, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 49, 25, 0, time.FixedZone("CET", 2*60*60)),
		File:      "cache.go",
		Message:   "Cache created",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 50, 31, 0, time.FixedZone("CET", 2*60*60)),
		File:      "rubberDuckDebugger.go",
		Message:   "Rubber duck debugging started",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 51, 9, 0, time.FixedZone("CET", 2*60*60)),
		File:      "server.go",
		Message:   "Error: Server is not responding",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 51, 33, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 51, 38, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 51, 45, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 51, 49, 0, time.FixedZone("CET", 2*60*60)),
		File:      "context.go",
		Message:   "Error: Failed to create context",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 52, 56, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 53, 10, 0, time.FixedZone("CET", 2*60*60)),
		File:      "theMatrix.go",
		Message:   "Error: There is no spoon",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 53, 32, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 54, 34, 0, time.FixedZone("CET", 2*60*60)),
		File:      "network.go",
		Message:   "Network connection established",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 57, 0, 0, time.FixedZone("CET", 2*60*60)),
		File:      "starTrek.go",
		Message:   "Error: Failed to beam up",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 57, 47, 0, time.FixedZone("CET", 2*60*60)),
		File:      "authentication.go",
		Message:   "Error: Failed to authenticate user",
	},
	{
		Timestamp: time.Date(2019, 4, 30, 12, 59, 49, 0, time.FixedZone("CET", 2*60*60)),
		File:      "memeGenerator.go",
		Message:   "Error: Meme generator ran out of memes",
	},
}
