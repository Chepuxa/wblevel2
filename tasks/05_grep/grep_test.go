package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"testing"
)

var inputFileName = "input.txt"

func TestGrep(t *testing.T) {
	wd, _ := os.Getwd()

	goCmdArgs := [][]string{
		{"run", "grep.go", "leaves", inputFileName},
		{"run", "grep.go", "n$", inputFileName},
		{"run", "grep.go", "-A=2", "n$", inputFileName},
		{"run", "grep.go", "-B=2", "n$", inputFileName},
		{"run", "grep.go", "-A=1", "-B=1", "n$", inputFileName},
		{"run", "grep.go", "-A=1", "-B=1", "-C=2", "n$", inputFileName},
		{"run", "grep.go", "-c", "n$", inputFileName},
		{"run", "grep.go", "-i", "N$", inputFileName},
		{"run", "grep.go", "-v", "n$", inputFileName},
		{"run", "grep.go", "-F", "-n", "To make a carpet on the ground.", inputFileName},
		{"run", "grep.go", "-c", "n$", inputFileName},
		{"run", "grep.go", "-c", "n$", inputFileName},
	}

	grepCmdArgs := [][]string{
		{"leaves", inputFileName},
		{"n$", inputFileName},
		{"-A2", "n$", inputFileName},
		{"-B2", "n$", inputFileName},
		{"-A1", "-B1", "n$", inputFileName},
		{"-A1", "-B1", "-C2", "n$", inputFileName},
		{"-c", "n$", inputFileName},
		{"-i", "N$", inputFileName},
		{"-v", "n$", inputFileName},
		{"-F", "-n", "To make a carpet on the ground.", inputFileName},
		{"-c", "n$", inputFileName},
		{"-c", "n$", inputFileName},
	}

	for i := range goCmdArgs {
		tcase := fmt.Sprintf("tcase%v", i+1)
		err := os.MkdirAll(tcase, 0777)
		if err != nil {
			t.Fatalf("Ошибка создания директории для тесткейса: %v", err)
		}

		filePathExpected := path.Join(tcase, "expected_output.txt")
		fileExpected, err := os.Create(filePathExpected)
		if err != nil {
			t.Fatalf("Ошибка создания файла: %v", err)
		}
		defer fileExpected.Close()

		filePathActual := path.Join(tcase, "output.txt")
		fileActual, err := os.Create(filePathActual)
		if err != nil {
			t.Fatalf("Ошибка создания файла: %v", err)
		}
		defer fileActual.Close()

		cmdGo := exec.Command("go", goCmdArgs[i]...)
		cmdGo.Dir = wd
		cmdGo.Stdout = fileActual
		cmdGo.Run()

		cmdGrep := exec.Command("grep", grepCmdArgs[i]...)
		cmdGrep.Dir = wd
		cmdGrep.Stdout = fileExpected
		cmdGrep.Run()

		fileExpected, err = os.Open(filePathExpected)
		if err != nil {
			t.Fatalf("Ошибка открытия файла: %v", err)
		}

		fileActual, err = os.Open(filePathActual)
		if err != nil {
			t.Fatalf("Ошибка открытия файла: %v", err)
		}

		expected, _ := io.ReadAll(fileExpected)
		result, _ := io.ReadAll(fileActual)

		if !bytes.Equal(result, expected) {
			t.Errorf("\nРасхождение результата в тесте %v, параметры: %v\n", tcase, goCmdArgs[i])
		}
	}
}
