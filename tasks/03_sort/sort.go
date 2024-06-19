package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var errNoFiles = errors.New("file names not specified")

func main() {
	flagK := flag.Int("k", 0, "Указание колонки для сортировки (по умолчанию 0 - вся строка)")
	flagN := flag.Bool("n", false, "Сортировать по числовому значению")
	flagR := flag.Bool("r", false, "Сортировать в обратном порядке")
	flagU := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flagF := flag.String("f", "", "Названия файлов через ,")
	flagO := flag.String("o", "output.txt", "Название файла для вывода")
	flag.Parse()

	if *flagF == "" {
		fmt.Println(errNoFiles)
		os.Exit(1)
	}

	filenames := strings.Split(*flagF, ",")

	lines := []string{}

	for _, v := range filenames {
		file, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmp := func(a, b string) int {
		ac := getColumn(a, *flagK)
		bc := getColumn(b, *flagK)

		var result int

		if *flagN {
			an, aerr := strconv.Atoi(ac)
			bn, berr := strconv.Atoi(bc)
			if aerr == nil && berr == nil {
				switch {
				case an > bn:
					result = 1
				case an < bn:
					result = -1
				}
			}
		} else {
			result = strings.Compare(ac, bc)
		}

		if *flagR {
			result *= -1
		}

		return result
	}

	slices.SortStableFunc(lines, cmp)

	if *flagU {
		lines = slices.Compact(lines)
	}

	f, err := os.Create(*flagO)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	for i, line := range lines {
		var toWrite string
		if i == 0 {
			toWrite = line
		} else {
			toWrite = "\n" + line
		}
		_, err := f.WriteString(toWrite)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func getColumn(s string, c int) string {
	f := strings.Fields(s)
	if c > 0 && c <= len(f) {
		return f[c-1]
	}

	return s
}
