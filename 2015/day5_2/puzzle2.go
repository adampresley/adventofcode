/*
--- Part Two ---

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
For example:

* qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
* xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
* uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
* ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.
How many strings are nice under these new rules?
*/

package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	log.Println("AdventOfCode.com - Day 5 - Puzzle 2")

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
