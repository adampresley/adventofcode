package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading puzzle file: %s", err.Error())
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	stringChannel := make(chan string, 100)
	statusChannel := make(chan StringStatus, 100)
	waitGroup := &sync.WaitGroup{}
	done := make(chan bool)

	niceCount := 0
	naughtyCount := 0

	go func() {
		for {
			select {
			case <-done:
				break

			case status := <-statusChannel:
				if status == NAUGHTY {
					naughtyCount++
				}

				if status == NICE {
					niceCount++
				}

				waitGroup.Done()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				break

			case stringToParse := <-stringChannel:
				stringParser := NewStringParser()
				status := stringParser.DetermineStatus(stringToParse)
				statusChannel <- status
			}
		}
	}()

	for reader.Scan() {
		waitGroup.Add(1)
		stringToParse := reader.Text()
		stringChannel <- stringToParse
	}

	waitGroup.Wait()
	done <- true

	log.Printf("%d are nice, %d are naughty\n", niceCount, naughtyCount)
}
