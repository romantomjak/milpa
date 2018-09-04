package main

import (
    "testing"
)

func Test_Maps_Letter_To_Code(t *testing.T) {
    letter := "R"
    code := "Romeo"
    result := LetterToCode(letter)
    if code != result {
        t.Errorf("Expected '%s' to be '%s', but got '%s'", letter, code, result)
    }
}

