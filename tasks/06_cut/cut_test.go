package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	goCmdArgs := [][]string{
		{"run", "cut.go", "-f=2,5"},
		{"run", "cut.go", "-f=2,5", "-d=;"},
		{"run", "cut.go", "-f=2,5", "-d=;", "-s"},
	}
	var stdout bytes.Buffer
	income := []string{
		"1 1Penalty for mr. Russel!",
		"2; 2Penalty; for; mr.; Russel!",
		"3;3Penalty;for;mr.;Russel!",
		"4-4Penalty-for-mr.-Russel!",
	}
	expected := []string{
		"1Penalty Russel!\n",
		" 2Penalty; Russel!\n",
		"3Penalty;Russel!\n",
		"\n",
	}

	for i := range goCmdArgs {
		cmd := exec.Command("go", goCmdArgs[i]...)
		cmd.Stdin = strings.NewReader(income[i])
		stdout.Reset()
		cmd.Stdout = &stdout

		err := cmd.Run()
		if err != nil {
			t.Fatalf("Ошибка выполнения команды: %v", err)
		}

		if result := stdout.String(); result != expected[i] {
			t.Errorf("Неожиданный результат №%v:\nфактический = %v\nожидаемый = %v\nаргументы - %v", i, result, expected[i], goCmdArgs[i])
		}
	}
}