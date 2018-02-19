package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// WordsCounter help structure
type WordsCounter struct {
	words int64
	chars int64
}

func main() {

	var lines []string
	var sumWords int64
	var sumChars int64
	var wg sync.WaitGroup

	start := time.Now()

	file, err := os.Open("..\\text.txt")
	if err != nil {
		fmt.Printf("error opening file %s", err)
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

	workers := len(lines)
	result := make(chan WordsCounter, workers)

	wg.Add(workers)

	for _, line := range lines {
		go worker(line, result, &wg)
	}

	wg.Wait()
	close(result)

	for wc := range result {
		sumWords += wc.words
		sumChars += wc.chars
	}

	timeElapsed := time.Since(start)
	fmt.Printf("Read %d lines, %d words %d chars in %fms\n", len(lines), sumWords, sumChars, timeElapsed.Seconds()*1000)
}

func worker(line string, result chan WordsCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	wordsCount := WordsCounter{}
	allwords := strings.Split(line, " ")
	var words []string
	for _, word := range allwords {
		if len(word) > 0 {
			words = append(words, word)
			wordsCount.chars += int64(len(word))
		}
	}
	wordsCount.words = int64(len(words))
	result <- wordsCount
}
