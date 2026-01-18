package main

import (
	"fmt"
	"os"
	_ "json_parser/tokens"
	"json_parser/scanner"
); 

func readJsonFile() []byte{
	// Here "data" is a type of []byte slice
	data, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	return data
	/**
	I learned that both []byte and []uint8 types are similar
	[]byte is an type alias of []uint8
	*/
}

func main() {
	var data_bytes []byte = readJsonFile()
	tokens := scanner.NewScanner(data_bytes)
	tokens.Tokenize()
}