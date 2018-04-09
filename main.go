package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	content := loadFile(fileName)

	tokenizer := Tokenizer{content: content, fileName: fileName}
	tokenizer.Scan()

	fmt.Print("TOKENS: ")
	fmt.Println(tokenizer.tokens)

	parser := NewParser(tokenizer.tokens)
	parser.start()
}

func loadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return content
}
