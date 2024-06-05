package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Зчитуємо текст з файлу
func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

// Індексуємо текст по словам (case insensitive)
func indexText(lines []string) map[string]struct{} {
	index := make(map[string]struct{})
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			lowerWord := strings.ToLower(word) // Перетворюємо слово на нижній регістр
			index[lowerWord] = struct{}{}
		}
	}
	return index
}

// Знаходимо всі рядки за словом (case insensitive)
func searchByWord(lines []string, query string) []string {
	lowerQuery := strings.ToLower(query) // Перетворюємо запит на нижній регістр
	var results []string
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), lowerQuery) {
			results = append(results, line)
		}
	}
	return results
}

// Пошук тексту
func searchText(lines []string) {
	fmt.Print("Введіть слово для пошуку: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	query = strings.TrimSpace(query) // Видаляємо символ нового рядка і пробіли

	results := searchByWord(lines, query)
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
	filename := "text.txt"
	lines := readFile(filename)

	// Індексуємо текст
	index := indexText(lines)
	fmt.Println("Проіндексовані слова (для тестування):")
	for word := range index {
		fmt.Printf("%s\n", word)
	}

	// Пошук тексту
	searchText(lines)
}
