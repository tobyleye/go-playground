package main

import (
	"fmt"
	"log"
	"os"
)

// reads the content of a file and spits it out

func readFile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("error reading file")
	}
	fileContent := string(file)
	fmt.Println(fileContent)
}

func readDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("error reading file", err)
	}
	for _, item := range dir {
		fmt.Println(item.Info())
		if item.IsDir() {
			fullPath := path + "/" + item.Name()
			readDir(fullPath)
		}
	}
}

func main() {
	// files := os.Args[1:]
	// fmt.Println("attempting to read the following files", files)
	// for _, file := range files {
	// 	readFile(file)
	// }
	dir := os.Args[2]
	fmt.Println("reading dir", dir)
	readDir(dir)
}
