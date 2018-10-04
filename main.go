package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var deleteVal bool
var searchVal string

func main() {
	searchPath := flag.String("p", ".", "Path to search")
	deleteFile := flag.Bool("d", false, "Delete the content")
	searchString := flag.String("str", "", "String to search inside files")
	fileName := flag.String("name", "", "Search for a specific file name")
	//replaceString := flag.String("r", "", "String that will replace the searched string")
	help := flag.Bool("h", false, "Print the help")
	folder := flag.Bool("f", false, "Print folder and files")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	} else {
		deleteVal = *deleteFile
		searchVal = *searchString
		filesRes := launchListing(*searchPath, *folder, *fileName)
		fileTreat(filesRes)
	}
}

func fileTreat(res []string) {
	for _, res := range files {
		if searchVal != "" {
			b, err := ioutil.ReadFile(res)
			if err != nil {
				panic(err)
			}
			s := string(b)

			if strings.Contains(s, searchVal) {
				fmt.Println(res)
			}
		} else {
			fmt.Println(res)
		}
	}
}
