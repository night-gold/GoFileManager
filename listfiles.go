package main

import (
	"os"
	"path/filepath"
	"strings"
)

var files []string
var folderVal bool
var fileNameVal string

func launchListing(path string, folder bool, fileName string) []string {
	folderVal = folder
	fileNameVal = fileName

	err := filepath.Walk(path, listFiles)

	if err != nil {
		panic(err)
	}

	return files
}

func listFiles(path string, info os.FileInfo, err error) error {
	if path != "." {
		if folderVal {
			if fileNameVal != "" {
				if strings.Contains(path, fileNameVal) {
					files = append(files, path)
				}
			} else {
				files = append(files, path)
			}
		} else if !info.IsDir() {
			if fileNameVal != "" {
				if strings.Contains(path, fileNameVal) {
					files = append(files, path)
				}
			} else {
				files = append(files, path)
			}
		}
	}
	return nil
}
