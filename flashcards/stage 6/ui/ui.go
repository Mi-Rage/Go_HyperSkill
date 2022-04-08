package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ui struct {
	LogStorage []string
}

func (u *Ui) AskTotalCards() int {
	var result int
	fmt.Println("Input the number of cards:")
	fmt.Scanln(&result)
	return result
}

func (u *Ui) GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	u.LogStorage = append(u.LogStorage, line)
	return line
}

func (u *Ui) GetCommand() string {
	commandSet := map[string]struct{}{
		"add":          {},
		"remove":       {},
		"import":       {},
		"export":       {},
		"ask":          {},
		"log":          {},
		"hardest card": {},
		"reset stats":  {},
	}
	var command string
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		command = u.GetUserInput()
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

func (u *Ui) LogAndPrint(data string) {
	u.LogStorage = append(u.LogStorage, data)
	fmt.Println(data)
}
