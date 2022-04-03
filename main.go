package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	dirPath := "./sample/"
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	newFilePrefix := "new_"
	fileExtension := ".txt"
	for i, f := range files {
		fmt.Println(f.Name())
		os.Rename(dirPath +f.Name(), dirPath + newFilePrefix + strconv.Itoa(i) + fileExtension)
	}
}