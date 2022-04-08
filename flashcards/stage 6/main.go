package main

import (
	. "task/card"
	. "task/storage"
	. "task/ui"
)

func main() {
	var storage = Storage{Data: make([]Card, 0)}
	var ui Ui = Ui{LogStorage: make([]string, 0)}

	for {
		command := ui.GetCommand()
		storage.ExecuteCommand(command, &ui)
	}
}
