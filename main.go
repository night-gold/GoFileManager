package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var deleteVal bool
var deleteFil bool
var file = false
var searchVal string
var filesListing []string
var newString string

func main() {
	searchPath := flag.String("p", ".", "Path to search")
	deleteContent := flag.Bool("d", false, "Delete the content")
	deleteFile := flag.Bool("df", false, "Delete all files returned with specific file name")
	searchString := flag.String("str", "", "String to search inside files")
	fileName := flag.String("name", "", "Search for a specific file name")
	replaceString := flag.String("r", "", "String that will replace the searched string")
	help := flag.Bool("h", false, "Print the help")
	folder := flag.Bool("f", false, "Print folder and files")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	} else {
		deleteVal = *deleteContent
		deleteFil = *deleteFile
		searchVal = *searchString
		newString = *replaceString
		if *fileName != "" {
			file = true
		}
		filesRes := launchListing(*searchPath, *folder, *fileName)
		fileTreat(filesRes)
	}
}

func fileTreat(res []string) {
	if searchVal != "" {
		for _, res := range files {
			b, err := ioutil.ReadFile(res)
			if err != nil {
				panic(err)
			}
			s := string(b)

			if strings.Contains(s, searchVal) {
				filesListing = append(files, res)
			}
		}

		fileRewrite(filesListing, searchVal, newString)
	} else if deleteFil {
		var validation = Ask4confirm()
		if (file) && (validation) {
			for _, res := range files {
				err := os.Remove(res)

				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			fmt.Println("Please give a fileName to delete")
		}
	} else {
		for _, res := range files {
			fmt.Println(res)
		}
	}
}
