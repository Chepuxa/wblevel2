package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Shell on Go. Type 'quit' to quit.")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "quit" {
			break
		}

		pipeline := strings.Split(input, " | ")

		for _, pipe := range pipeline {
			args := strings.Fields(pipe)
			if len(args) == 0 {
				continue
			}

			switch args[0] {
			case "cd":
				if len(args) < 2 {
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}

					err = os.Chdir(home)
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}
				} else {
					err := os.Chdir(args[1])
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}
				}
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
				}

				fmt.Println(dir)
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
			case "kill":
				if len(args) < 2 {
					fmt.Println("kill: missing argument")
				} else {
					pid, err := strconv.Atoi(args[1])
					if err != nil {
						fmt.Println(err)
					}

					proc, err := os.FindProcess(pid)
					if err != nil {
						fmt.Fprintf(os.Stderr, "kill: %v\n", err)
					}

					err = proc.Kill()
					if err != nil {
						fmt.Fprintf(os.Stderr, "kill: %v\n", err)
					}
				}
			case "ps":
				cmd := exec.Command("ps")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(cmd.Stderr, "ps: %v\n", err)
				}
			default:
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v: %v\n", args[0], err)
				}
			}
		}
	}
}
