package filemanager

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func (m *FileManager) ExportData(data []Card) {
	file, err := os.Create(m.FileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, datum := range data {
		_, err := fmt.Fprintln(file, datum.Term+":"+datum.Definition) // append the additional line
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%d cards have been saved.\n", len(data))
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

		loadedCards = append(loadedCards, Card{
			Term:       cardItem[0],
			Definition: cardItem[1],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return loadedCards
}
