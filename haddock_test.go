package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestGetRandomNumberBetween(t *testing.T) {
	min := 0
	max := 10
	for i := 0; i < 100; i++ {
		result := GetRandomNumberBetween(min, max)
		if result < min || result > max {
			t.Errorf("random number generation failed, %d not between %d and %d", result, min, max)
		}
	}
}

func TestGetRandomWordWithLength(t *testing.T) {
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

	for i := 6; i <= 22; i++ {
		word := GetRandomWordWithLength(words, i)
		if len(word) != i {
			t.Errorf("expected word of length %d, got length %d", i, len(word))
		}
	}
}

func TestGeneratePassword(t *testing.T) {
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

	for length := 16; length <= 48; length++ {
		password := GeneratePassword(length)
		if len(password) != length {
			t.Errorf("expected password of length %d, got length %d", length, len(password))
		}
	}
}
