package main

import (
	"fmt"
	"os"

	"github.com/jffp113/comparator"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage %s <file_path> <file_path>\n", os.Args[0])
		os.Exit(1)
	}

	fpath1 := os.Args[1]
	fpath2 := os.Args[2]

	file1, err := os.Open(fpath1)

	if err != nil {
		fmt.Printf("Error while opening the file: %v", err)
		os.Exit(1)
	}

	file2, err := os.Open(fpath2)

	if err != nil {
		fmt.Printf("Error while opening the file: %v\n", err)
		os.Exit(1)
	}

	equals, err := comparator.CompareJSON(file1, file2)

	if err != nil {
		fmt.Printf("Error while comparing both files: %v\n", err)
		os.Exit(1)
	}

	if equals {
		fmt.Println("JSON files are equal")
	} else {
		fmt.Println("JSON files different")
	}
}
