package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func generatePasswaord(length int, includeNumbers, includeSymbols bool) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if includeNumbers {
		charset += "0123456789"
	}

	if includeSymbols {
		charset += "!@#$%^&*()_+{}[]"
	}

	password := make([]byte, length)

	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password)
}

var wordList []string

func loadWordList() error {
	data, err := ioutil.ReadFile("wordlist.txt")

	if err != nil {
		return err
	}

	wordList = strings.Split(string(data), "\n")
	return nil
}

func generatePassphrase(length int) string {
	var words []string

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))
		words = append(words, wordList[randomIndex.Int64()])
	}

	return strings.Join(words, "-")
}

func main() {
	length := flag.Int("length", 12, "length of password")
	includeNumbers := flag.Bool("numbers", false, "include numbers")
	includeSymbols := flag.Bool("symbols", false, "include symbols")
	usePassphrase := flag.Bool("passphrase", false, "use passphrase")
	phraseLength := flag.Int("phrase-length", 3, "length of passphrase")

	flag.Parse()

	if *usePassphrase {
		err := loadWordList()

		if err != nil {
			fmt.Println(err)
			return
		}

		password := generatePassphrase(*phraseLength)
		fmt.Println(password)
		return
	}

	password := generatePasswaord(*length, *includeNumbers, *includeSymbols)
	fmt.Println(password)
}
