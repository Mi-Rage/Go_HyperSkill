package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName string
	var content []string
	fmt.Scanln(&fileName)
	content = readContent(fileName)
	for _, word := range content {
		fmt.Println(word)
	}
}

func readContent(name string) []string {
	var result []string
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}


