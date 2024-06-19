package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestSort(t *testing.T) {
	goCmdArgs := [][]string{
		{"run", "sort.go", "-k=2", "-f=tcase1/input1.txt,tcase1/input2.txt", "-o=tcase1/output.txt"},
		{"run", "sort.go", "-u", "-f=tcase2/input1.txt,tcase2/input2.txt", "-o=tcase2/output.txt"},
		{"run", "sort.go", "-n", "-f=tcase3/input1.txt,tcase3/input2.txt", "-o=tcase3/output.txt"},
		{"run", "sort.go", "-n", "-r", "-u", "-f=tcase4/input1.txt,tcase4/input2.txt", "-o=tcase4/output.txt"},
	}

	for i := range goCmdArgs {
		cmd := exec.Command("go", goCmdArgs[i]...)
		cmd.CombinedOutput()

		tcase := fmt.Sprintf("tcase%v/", i+1)

		filename := tcase + "expected_output.txt"
		fileExpected, err := os.Open(filename)
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer fileExpected.Close()

		fileSorted, err := os.Open(tcase + "output.txt")
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer fileSorted.Close()

		expected, _ := io.ReadAll(fileExpected)
		result, _ := io.ReadAll(fileSorted)

		if !bytes.Equal(result, expected) {
			t.Errorf("\nФактический результат = %v\nОжидаемый результат = %v\nаргументы - %v", result, expected, goCmdArgs[i])
		}
	}
}