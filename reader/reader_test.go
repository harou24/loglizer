package reader_test

import (
	"bufio"
	"bytes"
	"log"
	"loglizer/reader"
	"strings"
	"testing"
)

func TestReadHourlyLogBatches(t *testing.T) {
	// Expecting to receive 3 batches from the mock data
	const EXPECTED_NB_BATCHES = 3

	// Loading mock data to be read
	scanner := bufio.NewScanner(strings.NewReader(MockData))

	// Create a channel to capture grouped log entries
	logEntriesChan := make(chan []string, EXPECTED_NB_BATCHES)

	// Redirect log output to prevent cluttering the test output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		// Reset log output to default
		log.SetOutput(nil)
	}()

	// Call the function in a goroutine since it sends data to a channel
	go reader.ReadHourlyLogBatches(scanner, logEntriesChan)

	// Create a slice to hold the results received from the channel
	var results [][]string

	// Collect results from the channel
	for group := range logEntriesChan {
		results = append(results, group)
		// Closing channel when data has been received
		if len(results) == EXPECTED_NB_BATCHES {
			close(logEntriesChan)
		}
	}

	// We should receive exactly 3 batches
	if len(results) != EXPECTED_NB_BATCHES {
		t.Errorf("Expected 3 batches of log entries, got %d", len(results))
	}

}

// The first 100 lines of the CSV file
var MockData = `2019-04-30T12:01:39+02:00,network.go,Network connection established
2019-04-30T12:01:42+02:00,db.go,Transaction failed
2019-04-30T12:02:03+02:00,tardis.go,TARDIS dematerializing
2019-04-30T12:02:44+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:05:26+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:05:56+02:00,procrastinationService.go,Error: Failed to procrastinate
2019-04-30T12:06:19+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:07:20+02:00,hal9000.go,Error: I know that you and Frank were planning to disconnect me
2019-04-30T12:08:03+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:08:08+02:00,context.go,Error: Failed to create context
2019-04-30T12:08:23+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:08:38+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:10:43+02:00,server.go,Error: Server is not responding
2019-04-30T12:11:41+02:00,cache.go,Cache created
2019-04-30T12:11:41+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:11:53+02:00,tardis.go,TARDIS dematerializing
2019-04-30T12:15:02+02:00,rubberDuckDebugger.go,Rubber duck debugging started
2019-04-30T12:15:03+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:16:52+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:16:53+02:00,jokeGenerator.go,Error: Failed to generate a joke
2019-04-30T12:16:53+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:17:52+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:18:04+02:00,server.go,Error: Server is not responding
2019-04-30T12:18:06+02:00,coffeeMachine.go,Coffee is ready
2019-04-30T12:18:17+02:00,hal9000.go,Error: I know that you and Frank were planning to disconnect me
2019-04-30T12:18:24+02:00,rubberDuckDebugger.go,Rubber duck debugging started
2019-04-30T12:19:47+02:00,starTrek.go,Error: Failed to beam up
2019-04-30T12:19:56+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:20:20+02:00,cache.go,Cache created
2019-04-30T12:20:25+02:00,network.go,Network connection established
2019-04-30T12:21:37+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:21:39+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:21:43+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:21:45+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:22:47+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:23:59+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:26:13+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:28:28+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:29:08+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:29:33+02:00,cache.go,Cache created
2019-04-30T12:32:15+02:00,tardis.go,TARDIS dematerializing
2019-04-30T12:32:17+02:00,hal9000.go,Error: I know that you and Frank were planning to disconnect me
2019-04-30T12:34:05+02:00,db.go,Transaction failed
2019-04-30T12:34:46+02:00,cache.go,Cache created
2019-04-30T12:34:46+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:34:57+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:35:07+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:36:12+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:37:46+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:38:58+02:00,db.go,Transaction failed
2019-04-30T12:39:46+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:40:12+02:00,server.go,Error: Server is not responding
2019-04-30T12:40:17+02:00,network.go,Network connection established
2019-04-30T12:41:03+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:42:06+02:00,hal9000.go,Error: I know that you and Frank were planning to disconnect me
2019-04-30T12:42:29+02:00,server.go,Error: Server is not responding
2019-04-30T12:44:07+02:00,client.go,Error: Failed to connect to client
2019-04-30T12:45:25+02:00,authorization.go,Error: User is not authorized
2019-04-30T12:45:53+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:45:56+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:46:48+02:00,server.go,Error: Server is not responding
2019-04-30T12:47:35+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:47:49+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:48:01+02:00,network.go,Network connection established
2019-04-30T12:48:02+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:48:36+02:00,server.go,Error: Server is not responding
2019-04-30T12:49:25+02:00,cache.go,Cache created
2019-04-30T12:50:31+02:00,rubberDuckDebugger.go,Rubber duck debugging started
2019-04-30T12:51:09+02:00,server.go,Error: Server is not responding
2019-04-30T12:51:33+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:51:38+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:51:45+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:51:49+02:00,context.go,Error: Failed to create context
2019-04-30T12:52:56+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:53:10+02:00,theMatrix.go,Error: There is no spoon
2019-04-30T12:53:32+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T12:54:34+02:00,network.go,Network connection established
2019-04-30T12:57:00+02:00,starTrek.go,Error: Failed to beam up
2019-04-30T12:57:47+02:00,authentication.go,Error: Failed to authenticate user
2019-04-30T12:59:49+02:00,memeGenerator.go,Error: Meme generator ran out of memes
2019-04-30T13:01:41+02:00,hal9000.go,I'm sorry	 Dave. I'm afraid I can't do that
2019-04-30T13:03:54+02:00,rubberDuckDebugger.go,Error: Rubber duck went for a swim
2019-04-30T13:07:15+02:00,db.go,Error: Failed to connect to database
2019-04-30T13:12:16+02:00,coffeeMachine.go,Coffee is ready
2019-04-30T13:17:32+02:00,main.go,Listening on port 8080
2019-04-30T13:20:15+02:00,hal9000.go,I'm sorry	 Dave. I'm afraid I can't do that
2019-04-30T13:24:06+02:00,memeGenerator.go,Error: Failed to generate a meme
2019-04-30T13:26:21+02:00,main.go,Listening on port 8080
2019-04-30T13:27:24+02:00,client.go,Connected to client
2019-04-30T13:34:18+02:00,memeGenerator.go,Error: Failed to generate a meme
2019-04-30T13:37:37+02:00,memeGenerator.go,Error: Failed to generate a meme
2019-04-30T13:38:20+02:00,coffeeMachine.go,Error: Out of coffee beans
2019-04-30T13:38:38+02:00,context.go,Context created
2019-04-30T13:40:03+02:00,theMatrix.go,Red pill taken	 welcome to the real world
2019-04-30T13:51:37+02:00,coffeeMachine.go,Error: Out of coffee beans
2019-04-30T13:53:18+02:00,hal9000.go,Error: This mission is too important for me to allow you to jeopardize it
2019-04-30T13:53:56+02:00,db.go,Error: Failed to connect to database
2019-04-30T13:56:56+02:00,db.go,Transaction failed
2019-04-30T15:00:18+02:00,starTrek.go,Error: Failed to beam up
2019-04-30T15:00:28+02:00,starTrek.go,Error: Failed to beam up`
