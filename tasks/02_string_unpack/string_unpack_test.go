package main

import (
	"errors"
	"testing"
)

type TestData struct {
	input          string
	expectedOutput string
	expectedError  error
}

func TestUnpack(t *testing.T) {
	testData := []TestData{
		{"a2b2c2e", "aabbcce", nil},
		{"a2b23c2e", "", errDigitAfterDigit},
		{"\\a\\b\\c", "abc", nil},
		{"\\4\\2\\5", "425", nil},
		{"\\42", "44", nil},
		{"4a2a2", "", errDigitInBeginning},
		{"\\\\\\\\", "\\\\", nil},
		{"\\\\\\", "", errEscapeAtEnd},
		{"a0e0f4", "ffff", nil},
		{"\\40", "", nil},
		{"\\\\3", "\\\\\\", nil},
	}

	for _, v := range testData {
		res, err := unpackString(v.input)

		if res != v.expectedOutput {
			t.Errorf("\nОжидаемый результат: %v\nФактический результат: %v\n",
			v.expectedOutput, res)
		}

		if !errors.Is(err, v.expectedError) {
			t.Errorf("\nОжидаемая ошибка: %v\nФактическая ошибка: %v\n",
			v.expectedError, err)
		}
	}
}
