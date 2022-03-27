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
		censorship(getUserInput())
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

func censoredOutput(someWord string) {
	fmt.Println(wordProcessing(someWord))
}

func wordProcessing(someWord string) string {
	if isBadWord(someWord) {
		return strings.Repeat("*", len(someWord))
	}
	return someWord
}

func isBadWord(someWord string) bool {
	_, ok := badContent[strings.ToLower(someWord)]
	return ok
}

func censorship(someWord string) {
	switch someWord {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		censoredOutput(someWord)
	}
}



