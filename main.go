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

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
	}

	log.Println(GeneratePassword(words, 32))
}

func GeneratePassword(words map[int][]string, length int) string {
	wordone := GetRandomWord(words)
	wordtwo := GetRandomWordWithLength(words, length-len(wordone)-2)
	finalpassword := wordone + GetRandomDigit() + GetRandomSymbol() + wordtwo
	return finalpassword
}

func GetRandomWord(words map[int][]string) string {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		log.Fatal(err)
	}
	wordArray := words[int(i.Int64())]
	i, err = rand.Int(rand.Reader, big.NewInt(int64(len(wordArray))))
	if err != nil {
		log.Fatal(err)
	}
	return wordArray[int(i.Int64())]
}
func GetRandomWordWithLength(words map[int][]string, length int) string {
	wordArray := words[length]
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
