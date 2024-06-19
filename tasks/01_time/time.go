package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

/*

Создать программу печатающую точное время с использованием NTP-библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

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