package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fieldsToSelect := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", " ", "использовать другой разделитель")
	isSeparated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	if *fieldsToSelect == "" && !*isSeparated {
		fmt.Println("Необходимо указать хотя бы один из флагов: -f или -s")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *isSeparated || strings.Contains(line, *delimiter) {
			fields := strings.Split(line, *delimiter)

			if *fieldsToSelect != "" {
				selectedFields := selectFields(fields, *fieldsToSelect)
				fmt.Println(strings.Join(selectedFields, *delimiter))
			} else {
				fmt.Println(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения данных:", err)
		os.Exit(1)
	}
}

func selectFields(fields []string, fieldsList string) []string {
	selectedFields := []string{}
	fieldsIdx := parseFieldsList(fieldsList)

	for _, idx := range fieldsIdx {
		if idx > 0 && idx <= len(fields) {
			selectedFields = append(selectedFields, fields[idx-1])
		} else {
			selectedFields = append(selectedFields, "")
		}
	}

	return selectedFields
}

func parseFieldsList(fieldsList string) []int {
	fieldsIdx := []int{}
	fields := strings.Split(fieldsList, ",")

	for _, field := range fields {
		i := 0
		if n, err := fmt.Sscanf(field, "%d", &i); err == nil && n > 0 {
			fieldsIdx = append(fieldsIdx, i)
		}
	}

	return fieldsIdx
}
