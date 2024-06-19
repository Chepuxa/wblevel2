package main

import (
	"testing"
)

type TestData struct {
	input []string
	expectedOutput map[string]*[]string
}

func TestSort(t *testing.T) {
	testData := []TestData{
		{[]string{"аабб", "ббаа"}, map[string]*[]string{"аабб": {"ббаа", "аабб"}}},
		{[]string{"аабб", "ббаа", "ддд"}, map[string]*[]string{"аабб": {"ббаа", "аабб"}}},
		{[]string{"аАбБ", "БбАа"}, map[string]*[]string{"аабб": {"ббаа", "аабб"}}},
	}

	for _, data := range testData {
		res := searchAnagram(&data.input)
		for resultKey, resultValue := range *res {
			expectedWords, ok := data.expectedOutput[resultKey]
			if ok {
				for resultWordIdx, resultWord := range *resultValue {
					expectedWord := (*expectedWords)[resultWordIdx]
					if expectedWord != resultWord {
						t.Errorf("\nФактический результат = %v\nОжидаемый результат = %v\n", resultWord, expectedWord)
					}
				}
			} else {
				t.Errorf("\nФактический результат = %v\nОжидаемый результат = %v\n", resultValue, expectedWords)
			}
		}
	}
}

