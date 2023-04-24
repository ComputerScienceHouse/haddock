package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

var words map[int][]string
var longestWord int

const wordsFile = "./words.txt"

func main() {
	file, err := os.Open(wordsFile)
	if err != nil {
		log.Fatalf("Faild to open words file %s: %v", wordsFile, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read words file %s: %v", wordsFile, err)
	}

	longestWord = 0
	words = make(map[int][]string)

	for _, word := range lines {
		words[len(word)] = append(words[len(word)], word)
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}

	log.Printf("Successfully parsed words file %s", wordsFile)

	fileServer := http.FileServer(FileSystem{http.Dir("./static/")})
	http.HandleFunc("/api/v1/haddock", handleGeneratePassword)
	http.HandleFunc("/api/v1/xkcd", handleGenerateXKCDPassword)
	http.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fileServer))

	log.Println("Starting webserver on port 8000")

	err = http.ListenAndServe(":8000", nil)

	log.Fatalf("Webserver exited: %v", err)
}

func handleGeneratePassword(w http.ResponseWriter, r *http.Request) {
	var length int
	query := r.URL.Query()
	lenstr, present := query["length"]
	if !present || len(lenstr) == 0 {
		length = 24
	} else {
		lenth, err := strconv.Atoi(lenstr[0])
		if err != nil {
			length = 24
		}
		length = lenth
	}

	if length < 16 {
		length = 16
	} else if length > 48 {
		length = 48
	}

	passwordCount := 10

	data := make([]string, passwordCount)

	for i := range data {
		data[i] = GeneratePassword(length)
	}

	json.NewEncoder(w).Encode(data)
}

func handleGenerateXKCDPassword(w http.ResponseWriter, r *http.Request) {
	var length int
	query := r.URL.Query()
	lenstr, present := query["length"]
	if !present || len(lenstr) == 0 {
		length = 24
	} else {
		lenth, err := strconv.Atoi(lenstr[0])
		if err != nil {
			length = 24
		}
		length = lenth
	}

	if length < 16 {
		length = 16
	} else if length > 64 {
		length = 64
	}

	passwordCount := 10

	data := make([]string, passwordCount)

	for i := range data {
		data[i] = GenerateXKCDPassword(length)
	}

	json.NewEncoder(w).Encode(data)
}

// length MUST be above 16
func GeneratePassword(length int) string {
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
	lenone := GetRandomNumberBetween(min-1, ((length)/2)-1)
	lentwo := GetRandomNumberBetween(min-1, ((length)/2)-1)
	wordone := GetRandomWordWithLength(words, lenone)
	wordtwo := GetRandomWordWithLength(words, lentwo)
	finalpassword := wordone + GetRandomDigit()
	for len(finalpassword) < (length - lentwo - 1) {
		finalpassword = finalpassword + GetRandomDigit()
	}
	finalpassword = finalpassword + GetRandomSymbol() + wordtwo
	return finalpassword
}

func GenerateXKCDPassword(length int) string {
	minLen := (length-6)/4 - 2
	maxLen := (length-6)/4 + 2
	lengths := []int{GetRandomNumberBetween(minLen, maxLen), GetRandomNumberBetween(minLen, maxLen), GetRandomNumberBetween(minLen, maxLen)}
	lengths = append(lengths, length-3-(lengths[0]+lengths[1]+lengths[2]))
	return GetRandomWordWithLength(words, lengths[0]) + "-" + GetRandomWordWithLength(words, lengths[1]) + "-" + GetRandomWordWithLength(words, lengths[2]) + "-" + GetRandomWordWithLength(words, lengths[3])
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
		log.Fatal("Failed to generate random word with length %d: %v", length, err)
	}
	return wordArray[int(i.Int64())]
}

func GetRandomDigit() string {
	i, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to generate random digit: %v", err)
	}
	return strconv.Itoa(int(i.Int64()) % 10)
}

func GetRandomSymbol() string {
	symbols := []rune("`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?")
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
	if err != nil {
		log.Fatalf("Failed to generate random symbol: %v", err)
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
		log.Fatalf("Failed to generate random number: %v", err)
	}
	return int(i.Int64()) + 1
}

func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
