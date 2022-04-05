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
