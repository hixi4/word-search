package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func findLines(lines []string, query string) []string {
	var results []string
	for _, line := range lines {
		if strings.Contains(line, query) {
			results = append(results, line)
		}
	}
	return results
}

func searchText(lines []string) {
	fmt.Print("Введіть рядок для пошуку: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	query = strings.TrimSpace(query) // Убираем символ новой строки и пробелы

	results := findLines(lines, query)
	if len(results) == 0 {
		fmt.Println("Рядок не знайдено.")
		return
	}
	fmt.Println("Знайдені рядки:")
	for _, line := range results {
		fmt.Println(line)
	}
}

func main() {
	filename := "text.txt" // Замініть "your_file.txt" на ім'я вашого файлу
	lines := readFile(filename)
	searchText(lines)
}
