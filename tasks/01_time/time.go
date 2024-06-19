package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

var ntpServer = "0.beevik-ntp.pool.ntp.org"

func getTime() error {
	time, err := ntp.Time(ntpServer)

	if err != nil {
		return err
	}

	fmt.Println(time)

	return nil
}

func main() {
	if err := getTime(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}