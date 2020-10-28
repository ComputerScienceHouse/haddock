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

	log.Println(GeneratePassword(lines, 32))
}

func GeneratePassword(lines []string, length int) string {
	wordone := GetRandomWord(lines)
	wordtwo := GetRandomWord(lines)
	for len(wordtwo) != length-len(wordone)-2 {
		wordtwo = GetRandomWord(lines)
	}
	finalpassword := wordone + GetRandomDigit() + GetRandomSymbol() + wordtwo
	return finalpassword
}

func GetRandomWord(input []string) string {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(input))))
	if err != nil {
		log.Fatal(err)
	}
	return input[int(i.Int64())]
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
