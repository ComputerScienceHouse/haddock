package main

import (
	"bufio"
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
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

	log.Println("words.txt read")

	words := make(map[int][]string)
	longestWord := 0

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}

	log.Println(GeneratePassword(words, longestWord, 16))
	log.Println(GeneratePassword(words, longestWord, 16))
	log.Println(GeneratePassword(words, longestWord, 16))
	log.Println(GeneratePassword(words, longestWord, 24))
	log.Println(GeneratePassword(words, longestWord, 24))
	log.Println(GeneratePassword(words, longestWord, 48))
	log.Println(GeneratePassword(words, longestWord, 48))
	log.Println(GeneratePassword(words, longestWord, 48))
}

// length MUST be above 16
func GeneratePassword(words map[int][]string, longestWord int, length int) string {
	var min int
	if length <= 16 {
		min = 6
	} else if length <= 24 {
		min = 8
	} else if length <= 32 {
		min = 12
	} else {
		min = 14
	}
	lenone := GetRandomNumberBetween(min, ((length) / 2))
	lentwo := GetRandomNumberBetween(min, ((length) / 2))
	wordone := GetRandomWordWithLength(words, lenone)
	wordtwo := GetRandomWordWithLength(words, lentwo)
	finalpassword := wordone + GetRandomDigit()
	for len(finalpassword) < (length - lentwo - 1) {
		finalpassword = finalpassword + GetRandomDigit()
	}
	finalpassword = finalpassword + GetRandomSymbol() + wordtwo
	return finalpassword
}

func GetRandomWordWithLength(words map[int][]string, length int) string {
	var safelen int
	if length > 23 {
		safelen = 22
	} else {
		safelen = length
	}
	wordArray := words[safelen]
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordArray))))
	if err != nil {
		log.Fatal(err)
	}
	return wordArray[int(i.Int64())]
}

func GetRandomDigit() string {
	i, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(i.Int64()) % 10)
}

func GetRandomSymbol() string {
	// this is a somewhat restricted list of characters. some characters that may cause
	// problems in scripts have been removed. the original list is as follows:
	// `~!@#$%^&*()-_=+[{]}\|;:'",<.>/?
	symbols := []rune("!%^&*()-_=+[{]}|;:,<.>/?")
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
	if err != nil {
		log.Fatal(err)
	}
	return string(symbols[int(i.Int64())])
}

func GetRandomNumberBetween(min int, max int) int {
	num := (GetRandomNumber() % max) + 1
	for num < min || max < num {
		num = GetRandomNumber()
	}
	return num
}

func GetRandomNumber() int {
	i, err := rand.Int(rand.Reader, big.NewInt(63))
	if err != nil {
		log.Fatal(err)
	}
	return int(i.Int64()) + 1
}
