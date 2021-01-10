package main

import (
	"fmt"
	"io/ioutil"	
	"log"
	"strings"
	"os"
)

func main() {

	helloFilePath := os.Args[1]

	if helloFilePath == "" {
		log.Fatal("File path cannot be empty or nil")
	}

	helloGoText, err := updateHelloText(helloFilePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helloGoText)
}

func updateHelloText(filePath string) (string, error){
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	helloText := string(content)
	helloGoText := strings.Replace(helloText, "shell command", "golang program", 1)

	return helloGoText, nil
}