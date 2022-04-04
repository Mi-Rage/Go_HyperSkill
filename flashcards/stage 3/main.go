package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var totalCards int
	var storage = Storage{data: make([]Card, 0)}

	totalCards = askTotalCards()
	storage.initStorage(&totalCards)
	testUser(&storage)
}

func askTotalCards() int {
	var result int
	fmt.Println("Input the number of cards:")
	fmt.Scanln(&result)
	return result
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}

func testUser(s *Storage) {
	for _, card := range s.data {
		fmt.Printf("Print the definition of \"%s\":\n", card.Term)
		userAnswer := getUserInput()
		if card.Definition == userAnswer {
			fmt.Printf("Correct!\n")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\".\n", card.Definition)
		}
	}
}

type Card struct {
	Term       string
	Definition string
}

type Storage struct {
	data []Card
}

func (s *Storage) initStorage(cards *int) {
	var card Card
	var term, definition string
	for i := 0; i < *cards; i++ {
		fmt.Printf("The term for card #%d:\n", i+1)
		term = getUserInput()
		fmt.Printf("The definition for card #%d:\n", i+1)
		definition = getUserInput()
		card = Card{
			Term:       term,
			Definition: definition,
		}
		s.addCard(&card)
	}
}

func (s *Storage) addCard(c *Card) {
	s.data = append(s.data, *c)
}
