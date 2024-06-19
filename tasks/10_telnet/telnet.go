package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	flagTO := flag.String("timeout", "10s", "connection timeout")

	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go run task.go [--timeout=timeout] host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	timeOut, err := time.ParseDuration(*flagTO)
	if err != nil {
		fmt.Println("Error parsingg timeout duration:", err)
	}

	conn, err := net.DialTimeout("tcp", host+":"+port, timeOut)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	inputChan := make(chan string)

	go readFromSTDIN(inputChan)

	go readFromConnection(conn)

	for {
		select {
		case <-sigChan:
			fmt.Println("\nClosing connection...")

			return
		case input, ok := <-inputChan:
			if !ok {
				fmt.Println("\nClosing connection...2")

				return
			}

			input = strings.TrimSuffix(input, `\n`)
			input += "\r\nHost: " + host + `\r\n\r\n`

			_, err := conn.Write([]byte(input))
			if err != nil {
				fmt.Println("Error write to the connection:", err)
			}
		}
	}
}

func readFromSTDIN(inputChan chan string) {
	buf := make([]byte, 4096)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println(err)
			} else {
				fmt.Println("Error reading from STDIN:", err)
			}

			close(inputChan)

			return
		}

		input := string(buf[:n])

		inputChan <- input
	}
}

func readFromConnection(conn net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by remote host")
			} else {
				fmt.Println("Error reading from connection", err)
			}

			break
		}

		fmt.Println(string(buf[:n]))
	}

	os.Exit(0)
}
