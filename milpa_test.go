package main

import (
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

