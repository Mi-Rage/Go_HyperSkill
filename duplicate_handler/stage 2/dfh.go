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
	//Declaring and initializing the structure of the user interface
	var ui UserInterface
	ui.Directory = ui.GetDirectoryArg()
	ui.FileFormat = ui.GetFileFormat()
	ui.SortingOption = ui.GetSortOption()

	// Declaring the structure of the file finder and initialize the search results storage
	var finder = FileFinder{Storage: make(map[int][]FileObject)}
	// Searching for files by user-defined parameters
	finder.FindingFiles(ui.Directory, ui.FileFormat)

	//Output the search results in the user interface with user-defined parameters
	ui.SotredOutput(finder.Storage, ui.SortingOption)
}

// UserInterface - all communication with the user takes place in this structure
type UserInterface struct {
	Directory     string
	FileFormat    string
	SortingOption int8
}

// GetDirectoryArg - Getting the directory name from the command line
func (ui UserInterface) GetDirectoryArg() string {
	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		os.Exit(0)
	}
	return os.Args[1]
}

// GetFileFormat - Getting the directory name from the command line
func (ui UserInterface) GetFileFormat() string {
	var data string
	fmt.Println("Enter file format:")
	fmt.Scanln(&data)
	return "." + data
}

// GetSortOption - Getting the sorting order of the result
func (ui UserInterface) GetSortOption() int8 {
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

// SotredOutput - Output of search results taking into account the sorting order
func (ui UserInterface) SotredOutput(storage map[int][]FileObject, option int8) {
	// Extract all the sizes from the search results and put them in a slice
	var keys = make([]int, 0, len(storage))
	for key := range storage {
		keys = append(keys, key)
	}
	// Sort the resulting slice with the sizes
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
	// In the loop, for each size, extract the path from the storage and print it
	for _, key := range keys {
		fmt.Printf("%d bytes\n", key)
		for _, value := range storage[key] {
			fmt.Println(value.Path)
		}
	}
}

// FileObject - All search results are stored in the form of such structures
type FileObject struct {
	Path string
	Size int
}

// FileFinder - structure for searching files by specified values and saving the result to the repository
type FileFinder struct {
	Storage map[int][]FileObject
}

// FindingFiles - search for files in a given directory and save the result in a map
// with the key size and FileObject structure value
func (f FileFinder) FindingFiles(directory string, fileFormat string) {
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			if filepath.Ext(path) == fileFormat || fileFormat == "." {
				size := int(info.Size())
				f.Storage[size] = append(f.Storage[size], FileObject{Size: size, Path: path})
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
