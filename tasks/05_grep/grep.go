package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "Печатать +N строк после совпадения")
	before := flag.Int("B", 0, "Печатать +N строк до совпадения")
	context := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	isCount := flag.Bool("c", false, "Количество строк")
	isIgnoreCase := flag.Bool("i", false, "Игнорировать регистр")
	isInvert := flag.Bool("v", false, "Исключать вместо совпадения")
	isFixed := flag.Bool("F", false, "Точное совпадение со строкой вместо паттерна")
	isPrintLineNum := flag.Bool("n", false, "Напечатать номер строки")
	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	file, err := os.Open(path)
	if err != nil {
		raiseError(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		raiseError(err)
	}

	offsetUp, offsetDown := calculateOffsets(*after, *before, *context)

	var count int
	outputIndexes := []int{}
	result := make(map[int]string)

	for i, line := range lines {
		var isMatch bool

		if *isFixed {
			fixedLine := strings.TrimRight(line, "\r\n")
			if *isIgnoreCase {
				fixedLine = strings.ToLower(fixedLine)
			}
			isMatch = fixedLine == pattern
		} else {
			if *isIgnoreCase {
				pattern = "(?i)" + pattern
			}
			rg, err := regexp.Compile(pattern)
			if err != nil {
				raiseError(err)
			}
			isMatch = rg.MatchString(line)
		}

		if *isInvert {
			isMatch = !isMatch
		}

		if isMatch {
			count++

			for j := offsetUp; j > 0; j-- {
				iUp := i - j
				if iUp < 0 {
					break
				}
				if _, ok := result[iUp]; !ok {
					toAdd := lines[iUp]
					if *isPrintLineNum {
						toAdd = strconv.Itoa(iUp+1) + "-" + toAdd
					}
					result[iUp] = toAdd
					outputIndexes = append(outputIndexes, iUp)
				}
			}

			toAdd := line
			if *isPrintLineNum {
				toAdd = strconv.Itoa(i+1) + ":" + toAdd
			}

			if _, ok := result[i]; !ok {
				outputIndexes = append(outputIndexes, i)
			}
			result[i] = toAdd

			for j := 1; j <= offsetDown; j++ {
				iDown := i + j
				if iDown > len(lines)-1 {
					break
				}
				if _, ok := result[iDown]; !ok {
					toAdd := lines[iDown]
					if *isPrintLineNum {
						toAdd = strconv.Itoa(iDown+1) + "-" + toAdd
					}
					result[iDown] = toAdd
					outputIndexes = append(outputIndexes, iDown)
				}
			}
		}
	}

	if *isCount {
		fmt.Println(count)
	} else {
		for _, v := range outputIndexes {
			fmt.Println(result[v])
		}
	}
}

func raiseError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func calculateOffsets(after int, before int, context int) (offsetUp int, offsetDown int) {
	if context != 0 {
		offsetUp = context
		offsetDown = context
	}

	if after != 0 {
		offsetDown = after
	}

	if before != 0 {
		offsetUp = before
	}
	return offsetUp, offsetDown
}
