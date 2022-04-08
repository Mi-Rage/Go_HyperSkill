package main

import (
	"flag"
	. "task/card"
	. "task/storage"
	. "task/ui"
)

func main() {
	var storage = Storage{Data: make([]Card, 0)}
	var ui Ui = Ui{
		LogStorage: make([]string, 0),
	}

	importOpt := flag.String("import_from", "", "Enter file name to import")
	exportOpt := flag.String("export_to", "", "Enter file name to export")
	flag.Parse()
	ui.ImportOption = *importOpt
	ui.ExportOption = *exportOpt

	if ui.ImportOption != "" {
		storage.ImportCards(*importOpt, &ui)
	}

	for {
		command := ui.GetCommand()
		storage.ExecuteCommand(command, &ui)
	}
}


