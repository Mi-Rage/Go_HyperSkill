package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		os.Exit(0)
	}

	err := filepath.Walk(os.Args[1], func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
