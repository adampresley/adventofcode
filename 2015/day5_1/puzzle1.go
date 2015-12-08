/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
For example:

* ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
* aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
* jchzalrnumimnmhp is naughty because it has no double letter.
* haegwjzuvuyypxyu is naughty because it contains the string xy.
* dvszwmarrgswjxmb is naughty because it contains only one vowel.
How many strings are nice?
*/

package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	log.Println("AdventOfCode.com - Day 5 - Puzzle 1")

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
