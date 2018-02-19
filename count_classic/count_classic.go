package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var lines []string
	var sumWords int64
	var sumChars int64

	start := time.Now()

	file, err := os.Open("..\\text.txt")
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
		os.Exit(-1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		cWords, cChars := count(line)
		sumWords += cWords
		sumChars += cChars
	}

	timeElapsed := time.Since(start)
	fmt.Printf("Read %d lines, %d words %d chars in %fms\n", len(lines), sumWords, sumChars, timeElapsed.Seconds()*1000)
}

func count(line string) (countWords int64, countChars int64) {
	allwords := strings.Split(line, " ")
	var words []string
	for _, word := range allwords {
		if len(word) > 0 {
			words = append(words, word)
			countChars += int64(len(word))
		}
	}
	countWords = int64(len(words))
	return
}
