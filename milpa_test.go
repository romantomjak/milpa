package main

import (
    "testing"
)

var TEST_CODES = map[string]string{
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

func Test_Maps_Letters_To_Codes(t *testing.T) {
    for letter, code := range TEST_CODES {
        result := LetterToCode(letter)
        if code != result {
            t.Errorf("Expected '%s' to be a '%s' but got '%s'", letter, code, result)
        }
    }
}

