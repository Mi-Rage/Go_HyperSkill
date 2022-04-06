package storage

import (
	"fmt"
	. "task/card"
	. "task/filemanager"
	. "task/ui"
)

type Storage struct {
	Data []Card
}

func (s *Storage) ExecuteCommand(command string, ui *Ui) {
	switch command {
	case "add":
		s.createCard(ui)
	case "remove":
		s.removeCard()
	case "import":
		s.importCards()
	case "export":
		s.exportCards()
	case "ask":
		s.askUser(ui)
	}
}

func (s *Storage) createCard(ui *Ui) {
	var card Card
	var term, definition string

	fmt.Printf("The card :\n")
	for {
		term = ui.GetUserInput()
		if s.isUniqData(term, "") {
			break
		}
	}
	fmt.Printf("The definition of the card :\n")
	for {
		definition = ui.GetUserInput()
		if s.isUniqData("", definition) {
			break
		}
	}
	card = Card{
		Term:       term,
		Definition: definition,
	}
	s.addCard(&card)
	fmt.Printf("The pair(\"%s\":\"%s\") has been added.\n", term, definition)
}

func (s *Storage) addCard(c *Card) {
	s.Data = append(s.Data, *c)
}

func (s *Storage) isUniqData(t string, d string) bool {
	for _, datum := range s.Data {
		if len(t) != 0 {
			if datum.Term == t {
				fmt.Printf("The card \"%s\" already exists. Try again:\n", t)
				return false
			}
		} else {
			if datum.Definition == d {
				fmt.Printf("The definition \"%s\" already exists. Try again:\n", d)
				return false
			}
		}
	}
	return true
}

func (s *Storage) removeCard() {
	var cardToRemove string
	var isDeleted = false
	fmt.Println("Which card?")
	fmt.Scanln(&cardToRemove)
	for i, datum := range s.Data {
		if datum.Term == cardToRemove {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			isDeleted = true
			break
		}
	}
	if isDeleted {
		fmt.Println("The card has been removed.")
	} else {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n", cardToRemove)
	}
}

func (s *Storage) importCards() {
	var fileName string

	fmt.Println("File name:")
	fmt.Scanln(&fileName)
	fileManager := FileManager{FileName: fileName}
	if fileManager.IsFileExist() {
		loadedStrings := 0
		for _, card := range fileManager.ImportData() {
			isDoubleFound := false
			loadedStrings += 1
			for index, datum := range s.Data {
				if datum.Term == card.Term {
					s.Data[index].Definition = card.Definition
					isDoubleFound = true
					break
				}
			}
			if !isDoubleFound {
				s.Data = append(s.Data, card)
			}
		}
		fmt.Printf("%d cards have been loaded.\n", loadedStrings)
	} else {
		fmt.Println("File not found.")
	}
}

func (s *Storage) exportCards() {
	var fileName string
	fmt.Println("File name:")
	fmt.Scanln(&fileName)
	fileManager := FileManager{FileName: fileName}
	fileManager.ExportData(s.Data)
}

func (s *Storage) askUser(ui *Ui) {
	var timesToAsk int
	fmt.Println("How many times to ask?")
	fmt.Scanln(&timesToAsk)
	for i := 0; i < timesToAsk; i++ {
		cardTerm := s.Data[i%len(s.Data)].Term
		cardDef := s.Data[i%len(s.Data)].Definition
		fmt.Printf("Print the definition of \"%s\":\n", cardTerm)
		userAnswer := ui.GetUserInput()
		if cardDef == userAnswer {
			fmt.Printf("Correct!\n")
		} else {
			rightAnswer, foundTerm := s.checkWrongAnswer(userAnswer, cardTerm)
			if foundTerm == "" {
				fmt.Printf("Wrong. The right answer is \"%s\".\n", cardDef)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\", "+
					"but your definition is correct for \"%s\"\n", rightAnswer, foundTerm)
			}
		}
	}
}

func (s *Storage) checkWrongAnswer(answer string, currentTerm string) (string, string) {
	var rightDef, foundTrm string
	for _, datum := range s.Data {
		if answer == datum.Definition {
			foundTrm = datum.Term
			break
		}
	}
	for _, datum := range s.Data {
		if datum.Term == currentTerm {
			rightDef = datum.Definition
			break
		}
	}
	return rightDef, foundTrm
}
