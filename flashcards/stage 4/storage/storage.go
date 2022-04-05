package storage

import (
	"fmt"
	. "task/card"
	. "task/ui"
)

type Storage struct {
	Data []Card
}

func (s *Storage) InitStorage(cards *int, ui *Ui) {
	var card Card
	var term, definition string

	for i := 0; i < *cards; i++ {
		fmt.Printf("The term for card #%d:\n", i+1)
		for {
			term = ui.GetUserInput()
			if s.isUniqData(term, "") {
				break
			}
		}
		fmt.Printf("The definition for card #%d:\n", i+1)
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
	}
}

func (s *Storage) addCard(c *Card) {
	s.Data = append(s.Data, *c)
}

func (s *Storage) isUniqData(t string, d string) bool {
	for _, datum := range s.Data {
		if len(t) != 0 {
			if datum.Term == t {
				fmt.Printf("The term \"%s\" already exists. Try again:\n", t)
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

func (s *Storage) TestUser(ui *Ui) {
	for _, card := range s.Data {
		fmt.Printf("Print the definition of \"%s\":\n", card.Term)
		userAnswer := ui.GetUserInput()
		if card.Definition == userAnswer {
			fmt.Printf("Correct!\n")
		} else {
			rightAnswer, foundTerm := s.checkWrongAnswer(userAnswer, card.Term)
			if foundTerm == "" {
				fmt.Printf("Wrong. The right answer is \"%s\".\n", card.Definition)
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
