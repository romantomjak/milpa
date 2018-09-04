package main

import (
    "strings"
    "testing"
)

func Test_Maps_Letters_To_Codes(t *testing.T) {
    for letter, code := range CODES {
        result := LetterToCode(letter)
        if code != result {
            t.Errorf("Expected '%s' to be a '%s' but got '%s'", letter, code, result)
        }
    }
}

func Test_Ignores_Unknown_Symbols(t *testing.T) {
    symbols := []string{" ", ",", ";", "!"}
    for _, symbol := range symbols {
        result := LetterToCode(symbol)
        if symbol != result {
            t.Errorf("Expected '%s' to be the same, but got '%s'", symbol, result)
        }
    }
}

func Test_Ignores_Case(t *testing.T) {
    for letter, code := range CODES {
        lcLetter := strings.ToLower(letter)
        result := LetterToCode(lcLetter)
        if code != result {
            t.Errorf("Expected '%s' to be a '%s' but got '%s'", lcLetter, code, result)
        }
    }
}

func Test_Maps_Word_To_Codes(t *testing.T) {
    testCases := []struct {
        words string
        want  string
    }{
        {"Foo", "Foxtrot Oscar Oscar"},
        {"Foo Bar", "Foxtrot Oscar Oscar Bravo Alpha Romeo"},
    }
    for _, tc := range testCases {
        if got := WordToCode(tc.words); got != tc.want {
            t.Errorf("Expected '%s' to be a '%s' but got '%s'", tc.words, tc.want, got)
        }
    }
}
