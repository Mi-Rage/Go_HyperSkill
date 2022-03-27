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
	var checkedWord string

	fileName = getUserInput()
	checkedWord = strings.ToLower(getUserInput())

	readContent(fileName)
	_, ok := badContent[checkedWord]
	fmt.Println(ok)
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


