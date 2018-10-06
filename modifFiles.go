package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func fileRewrite(list []string, search string, replace string) {
	fmt.Println("The string was replace in these file : ")
	for _, list := range files {
		read, err := ioutil.ReadFile(list)
		if err != nil {
			panic(err)
		}
		newContent := strings.Replace(string(read), search, replace, -1)

		err = ioutil.WriteFile(list, []byte(newContent), 0)
		if err != nil {
			panic(err)
		}

		fmt.Println(list)
	}
}
