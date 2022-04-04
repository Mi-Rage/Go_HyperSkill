package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var card = Card{
		Term:       getUserInput(),
		Definition: getUserInput(),
	}
	answer := getUserInput()
	if card.Definition == answer {
		fmt.Println("right")
	} else {
		fmt.Println("wrong")
	}
}

type Card struct {
	Term       string
	Definition string
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}
