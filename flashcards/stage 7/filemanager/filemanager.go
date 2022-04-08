package filemanager

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	. "task/card"
)

type FileManager struct {
	FileName string
}

func (m *FileManager) IsFileExist() bool {
	_, err := os.Stat(m.FileName)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func (m *FileManager) ExportData(data []Card) int {
	file, err := os.Create(m.FileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, datum := range data {
		_, err := fmt.Fprintln(file, datum.Term+":"+datum.Definition+":"+strconv.Itoa(datum.Error)) // append the additional line
		if err != nil {
			log.Fatal(err)
		}
	}
	return len(data)
}

func (m *FileManager) ImportData() []Card {
	var loadedCards = make([]Card, 0)

	file, err := os.Open(m.FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // create a new Scanner for the file

	for scanner.Scan() {
		cardItem := strings.Split(scanner.Text(), ":")
		cardError, _ := strconv.Atoi(cardItem[2])

		loadedCards = append(loadedCards, Card{
			Term:       cardItem[0],
			Definition: cardItem[1],
			Error:      cardError,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return loadedCards
}

func (m *FileManager) ExportLogs(storage []string) {
	file, err := os.Create(m.FileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, datum := range storage {
		_, err := fmt.Fprintln(file, datum)
		if err != nil {
			log.Fatal(err)
		}
	}
}
