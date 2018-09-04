package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var CODES = map[string]string{
	"A": "Alpha",
	"B": "Bravo",
	"C": "Charlie",
	"D": "Delta",
	"E": "Echo",
	"F": "Foxtrot",
	"G": "Golf",
	"H": "Hotel",
	"I": "India",
	"J": "Juliett",
	"K": "Kilo",
	"L": "Lima",
	"M": "Mike",
	"N": "November",
	"O": "Oscar",
	"P": "Papa",
	"Q": "Quebec",
	"R": "Romeo",
	"S": "Sierra",
	"T": "Tango",
	"U": "Uniform",
	"V": "Victor",
	"W": "Whiskey",
	"X": "X-ray",
	"Y": "Yankee",
	"Z": "Zulu",
}

func LetterToCode(letter string) string {
	code := letter
	if strings.ToLower(letter) == letter {
		code = strings.ToUpper(letter)
	}
	if val, ok := CODES[code]; ok {
		return val
	}
	return code
}

func WordToCode(word string) string {
	var buffer bytes.Buffer
	for index, character := range word {
		letter := string(character)
		if letter == " " { // don't process spaces
			continue
		}
		code := LetterToCode(letter)
		space := " "
		if index+1 == len(word) { // skip trailing space
			space = ""
		}
		buffer.WriteString(code + space)
	}
	return buffer.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s hello world\n", os.Args[0])
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		fmt.Println(WordToCode(word))
	}
}
