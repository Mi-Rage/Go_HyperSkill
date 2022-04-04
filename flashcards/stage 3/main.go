package main

import (
	. "task/card"
	. "task/storage"
	. "task/ui"
)

func main() {
	var totalCards int
	var storage = Storage{Data: make([]Card, 0)}
	var ui Ui

	totalCards = ui.AskTotalCards()
	storage.InitStorage(&totalCards, &ui)
	storage.TestUser(&ui)
}
