package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ui struct {
}

func (u Ui) AskTotalCards() int {
	var result int
	fmt.Println("Input the number of cards:")
	fmt.Scanln(&result)
	return result
}

func (u Ui) GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return line
}

func (u Ui) GetCommand() string {
	commandSet := map[string]struct{}{
		"add":    {},
		"remove": {},
		"import": {},
		"export": {},
		"ask":    {},
	}
	var command string
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")
		fmt.Scanln(&command)
		if command == "exit" {
			fmt.Println("Bye bye!")
			os.Exit(0)
		}
		if _, ok := commandSet[command]; ok {
			return command
		} else {
			fmt.Println("Wrong command. Try again:")
		}
	}
}
