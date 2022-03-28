package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var badContent = make(map[string]struct{}, 3)

func main() {
	var fileName string

	fileName = getUserInput()
	readContent(fileName)
	for {
		censorship(getUserSentences())
	}
}

func getUserInput() string {
	var data string
	fmt.Scanln(&data)
	return data
}

func readContent(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		badContent[strings.ToLower(scanner.Text())] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getUserSentences() []string {
	var result []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result = strings.Split(scanner.Text(), " ")
	return result
}

func censoredOutput(someWords []string) {
	fmt.Println(wordProcessing(someWords))
}

func wordProcessing(someWords []string) string {
	var b strings.Builder
	for index, word := range someWords {
		if isBadWord(word) {
			b.WriteString(strings.Repeat("*", len(word)))
		} else {
			b.WriteString(word)
		}
		if index != len(someWords) - 1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func isBadWord(someWord string) bool {
	_, ok := badContent[strings.ToLower(someWord)]
	return ok
}

func censorship(someWords []string) {
	if len(someWords) == 1 && someWords[0] == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	censoredOutput(someWords)
}



