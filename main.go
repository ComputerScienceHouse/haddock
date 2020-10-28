package main

import (
	"bufio"
	"log"
	"math/rand"
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
	return input[rand.Intn(len(input))]
}

func GetRandomDigit() string {
	return strconv.Itoa(rand.Int() % 10)
}

func GetRandomSymbol() string {
	// this is a somewhat restricted list of characters. some characters that may cause
	// problems in scripts have been removed. the original list is as follows:
	// `~!@#$%^&*()-_=+[{]}\|;:'",<.>/?
	symbols := []rune("!%^&*()-_=+[{]}|;:,<.>/?")
	return string(symbols[rand.Intn(len(symbols))])

}
