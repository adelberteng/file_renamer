package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"strconv"
	"flag"
)

func main() {
	dir := flag.String("dir", "", "files in directory need to rename.")
	prefix := flag.String("prefix", "", "prefix of new file name.")
	ext := flag.String("ext", "", "file extension.")
	flag.Parse()

	if *dir == "" {
		log.Fatal("directory path is empty.")
	}

	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	filePrefix := *prefix + "_"
	for i, f := range files {
		fmt.Println(f.Name(), "to", filePrefix + strconv.Itoa(i) + "." + *ext)
		os.Rename(*dir +f.Name(), *dir + filePrefix + strconv.Itoa(i) + *ext)
		
	}
}