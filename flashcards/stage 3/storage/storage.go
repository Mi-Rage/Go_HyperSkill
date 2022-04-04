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
		term = ui.GetUserInput()
		fmt.Printf("The definition for card #%d:\n", i+1)
		definition = ui.GetUserInput()
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

func (s *Storage) TestUser(ui *Ui) {
	for _, card := range s.Data {
		fmt.Printf("Print the definition of \"%s\":\n", card.Term)
		userAnswer := ui.GetUserInput()
		if card.Definition == userAnswer {
			fmt.Printf("Correct!\n")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\".\n", card.Definition)
		}
	}
}
