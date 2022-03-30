package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {

	var directory, fileFormat string
	var sortOption int8
	var storage = make(map[int][]string)

	directory = getDirectoryArg()
	fileFormat = getFileFormat()
	sortOption = getSortOption()
	findFiles(directory, fileFormat, storage)
	sotredOutput(storage, sortOption)
}

func getDirectoryArg() string {
	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		os.Exit(0)
	}
	return os.Args[1]
}

func getFileFormat() string {
	var data string
	fmt.Println("Enter file format:")
	fmt.Scanln(&data)
	return "." + data
}

func getSortOption() int8 {
	fmt.Println("Size sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")
	var option int8
	for {
		fmt.Println("Enter a sorting option:")
		fmt.Scanln(&option)
		if option != 1 && option != 2 {
			fmt.Println("Wrong option")
		} else {
			return option
		}
	}
}

func findFiles(directory string, fileFormat string, storage map[int][]string) {
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			if filepath.Ext(path) == fileFormat || fileFormat == "." {
				size := info.Size()
				if storage[int(size)] != nil {
					storage[int(size)] = append(storage[int(size)], path)
				} else {
					storage[int(size)] = []string{path}
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func sotredOutput(storage map[int][]string, option int8) {
	var keys = make([]int, 0, len(storage))
	for k := range storage {
		keys = append(keys, k)
	}
	switch option {
	case 1:
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})
	case 2:
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
	}
	for _, k := range keys {
		fmt.Printf("%d bytes\n", k)
		for _, v := range storage[k] {
			fmt.Println(v)
		}
	}
}
