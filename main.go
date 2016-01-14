package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mantishK/quiz/tree"
)

var fName string

func init() {
	flag.StringVar(&fName, "file", "word.list", "File name")
	flag.Parse()
}

func main() {
	tc := tree.TreeCollection{}
	buildTree(&tc)
	w := longCompWord(&tc)
	fmt.Println("Longest compound word in the list: ", w)
}

func buildTree(tc *tree.TreeCollection) {
	file, err := os.Open(fName)
	if err != nil {
		panic("File not found")
	}
	s := bufio.NewScanner(file)
	for s.Scan() {
		word := s.Text()
		tc.AddWord(word)
	}
	defer file.Close()
	if err := s.Err(); err != nil {
		panic("Error scanning file")
	}
}

func longCompWord(tc *tree.TreeCollection) string {
	maxLength := 0
	longWord := ""
	file, err := os.Open(fName)
	defer file.Close()
	if err != nil {
		panic("File not found")
	}
	s := bufio.NewScanner(file)
	for s.Scan() {
		word := s.Text()
		if tc.IsCompound(word) && maxLength < len(word) {
			longWord = word
			maxLength = len(word)
		}
	}
	if err := s.Err(); err != nil {
		panic("Error scanning file")
	}
	return longWord
}
