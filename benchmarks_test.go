package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func BenchmarkGetRandomWordWithLength6(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetRandomWordWithLength(words, 6)
	}
}

func BenchmarkGetRandomWordWithLength16(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetRandomWordWithLength(words, 16)
	}
}

func BenchmarkGetRandomWordWithLength22(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetRandomWordWithLength(words, 22)
	}
}

func BenchmarkGeneratePasswordWithLength16(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GeneratePassword(16)
	}
}

func BenchmarkGeneratePasswordWithLength32(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GeneratePassword(32)
	}
}

func BenchmarkGeneratePasswordWithLength64(b *testing.B) {
	// ----------------------------
	//        BEGIN SETUP
	// ----------------------------
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	// ----------------------------
	//         END SETUP
	// ----------------------------
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GeneratePassword(64)
	}
}
