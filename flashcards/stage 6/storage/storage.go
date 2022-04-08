package storage

import (
	"fmt"
	"strconv"
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
		s.removeCard(ui)
	case "import":
		s.importCards(ui)
	case "export":
		s.exportCards(ui)
	case "ask":
		s.askUser(ui)
	case "log":
		s.logData(ui)
	case "hardest card":
		s.getHardestCard(ui)
	case "reset stats":
		s.setResetErrors(ui)
	}
}

func (s *Storage) createCard(ui *Ui) {
	var card Card
	var term, definition string

	ui.LogAndPrint("The card :")
	for {
		term = ui.GetUserInput()
		if s.isUniqData(term, "", ui) {
			break
		}
	}
	ui.LogAndPrint("The definition of the card :")
	for {
		definition = ui.GetUserInput()
		if s.isUniqData("", definition, ui) {
			break
		}
	}
	card = Card{
		Term:       term,
		Definition: definition,
		Error:      0,
	}
	s.addCard(&card)
	ui.LogAndPrint(fmt.Sprintf("The pair(\"%s\":\"%s\") has been added.", term, definition))
}

func (s *Storage) addCard(c *Card) {
	s.Data = append(s.Data, *c)
}

func (s *Storage) isUniqData(t string, d string, ui *Ui) bool {
	for _, datum := range s.Data {
		if len(t) != 0 {
			if datum.Term == t {
				ui.LogAndPrint(fmt.Sprintf("The card \"%s\" already exists. Try again:", t))
				return false
			}
		} else {
			if datum.Definition == d {
				ui.LogAndPrint(fmt.Sprintf("The definition \"%s\" already exists. Try again:", d))
				return false
			}
		}
	}
	return true
}

func (s *Storage) removeCard(ui *Ui) {
	var cardToRemove string
	var isDeleted = false

	ui.LogAndPrint("Which card?")
	cardToRemove = ui.GetUserInput()
	for i, datum := range s.Data {
		if datum.Term == cardToRemove {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			isDeleted = true
			break
		}
	}
	if isDeleted {
		ui.LogAndPrint("The card has been removed.")
	} else {
		ui.LogAndPrint(fmt.Sprintf("Can't remove \"%s\": there is no such card.", cardToRemove))
	}
}

func (s *Storage) importCards(ui *Ui) {
	var fileName string

	ui.LogAndPrint("File name:")
	fileName = ui.GetUserInput()
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
		ui.LogAndPrint(fmt.Sprintf("%d cards have been loaded.", loadedStrings))
	} else {
		ui.LogAndPrint("File not found.")
	}
}

func (s *Storage) exportCards(ui *Ui) {
	var fileName string

	ui.LogAndPrint("File name:")
	fileName = ui.GetUserInput()
	fileManager := FileManager{FileName: fileName}
	ui.LogAndPrint(fmt.Sprintf("%d cards have been saved.", fileManager.ExportData(s.Data)))
}

func (s *Storage) askUser(ui *Ui) {
	var timesToAsk int

	ui.LogAndPrint("How many times to ask?")
	timesToAsk, _ = strconv.Atoi(ui.GetUserInput())
	for i := 0; i < timesToAsk; i++ {
		cardTerm := s.Data[i%len(s.Data)].Term
		cardDef := s.Data[i%len(s.Data)].Definition
		ui.LogAndPrint(fmt.Sprintf("Print the definition of \"%s\":", cardTerm))
		userAnswer := ui.GetUserInput()
		if cardDef == userAnswer {
			ui.LogAndPrint("Correct!")
		} else {
			rightAnswer, foundTerm := s.checkWrongAnswer(userAnswer, cardTerm)
			if foundTerm == "" {
				ui.LogAndPrint(fmt.Sprintf("Wrong. The right answer is \"%s\".", cardDef))
			} else {
				ui.LogAndPrint(fmt.Sprintf("Wrong. The right answer is \"%s\", "+
					"but your definition is correct for \"%s\"", rightAnswer, foundTerm))
			}
			s.Data[i%len(s.Data)].Error += 1
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

func (s *Storage) logData(ui *Ui) {
	var fileName string

	ui.LogAndPrint("File name:")
	fileName = ui.GetUserInput()
	fileManager := FileManager{FileName: fileName}
	fileManager.ExportLogs(ui.LogStorage)
	ui.LogAndPrint("The log has been saved.")
}

func (s *Storage) getHardestCard(ui *Ui) {
	var hardestCards []string
	var maxError = 0

	for _, datum := range s.Data {
		if maxError < datum.Error {
			maxError = datum.Error
		}
	}

	if maxError > 0 {
		for _, datum := range s.Data {
			if datum.Error == maxError {
				hardestCards = append(hardestCards, datum.Term)
			}
		}
	}

	if len(hardestCards) == 0 {
		ui.LogAndPrint("There are no cards with errors.")
	} else if len(hardestCards) == 1 {
		ui.LogAndPrint(fmt.Sprintf("The hardest card is \"%s\". "+
			"You have %d errors answering it", hardestCards[0], maxError))
	} else {
		var ans = "The hardest card is "
		for _, card := range hardestCards {
			ans += "\"" + card + "\", "
		}
		ans = ans[:len(ans)-2] + fmt.Sprintf(". You have %d errors answering them.", maxError)
		ui.LogAndPrint(ans)
	}

}

func (s *Storage) setResetErrors(ui *Ui) {
	for i, _ := range s.Data {
		s.Data[i].Error = 0
	}
	ui.LogAndPrint("Card statistics have been reset.")
}
