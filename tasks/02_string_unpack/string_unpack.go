package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var str = "a2b2"
var errDigitAfterDigit = errors.New("invalid sequence: digit placed after digit")
var errDigitInBeginning = errors.New("invalid sequence: digit in beginning")
var errEscapeAtEnd = errors.New("invalid sequence: escape character at the end of the string")

func unpackString(str string) (string, error) {
	runes := []rune(str)
	var s strings.Builder
	var prevRune rune
	var escapeNext bool
	var prevIsUnescapedDigit bool
	for i, v := range runes {
		switch {
		case unicode.IsDigit(v) && !escapeNext:
			if prevIsUnescapedDigit {
				return "", errDigitAfterDigit
			}
			if prevIsUnescapedDigit || prevRune == 0 {
				return "", errDigitInBeginning
			}
			num, _ := strconv.Atoi(string(v))
			s.WriteString(strings.Repeat(string(prevRune), num))
			prevRune = v
			prevIsUnescapedDigit = true
		case v == '\\' && !escapeNext:
			if i == len(runes)-1 {
				return "", errEscapeAtEnd
			}
			escapeNext = true
		default:
			if prevRune != 0 && !prevIsUnescapedDigit {
				s.WriteRune(prevRune)
			}

			if i == len(runes)-1 {
				s.WriteRune(v)
				break
			}

			prevRune = v
			escapeNext = false
			prevIsUnescapedDigit = false
		}
	}
	return s.String(), nil
}

func main() {
	res, err := unpackString(str)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(res)
}
