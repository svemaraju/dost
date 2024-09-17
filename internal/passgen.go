package internal

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

var (
	noSymbols       bool
	copyToClipBoard bool = true
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
)

func writeToClipboard(text string) error {
	return clipboard.WriteAll(text)
}


func shuffle(password []string) []string {
	mr := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	mr.Shuffle(len(password), func(i, j int) { 
		password[i], password[j] = password[j], password[i]
	})
	return password
}

// selectRandomCharacter picks a random element from slice of strings
func selectRandomCharacter(characterSet string) string {
	numOfCharacters := len(characterSet)
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(numOfCharacters)))
	if err != nil {
		panic(err)
	}
	return string(characterSet[randomIndex.Int64()])
}

// generatePassword returns a random password of the specified length
func generatePassword(length int, noSymbols bool) (string, error) {
	// Combine all characters for password generation
	allChars := uppercaseLetters + lowercaseLetters + digits

	var password []string

	mandatoryUppercaseChar := selectRandomCharacter(uppercaseLetters)
	password = append(password, mandatoryUppercaseChar)

	mandatoryLowercaseChar := selectRandomCharacter(lowercaseLetters)
	password = append(password, mandatoryLowercaseChar)

	mandatorydigit := selectRandomCharacter(digits)
	password = append(password, mandatorydigit)

	if !noSymbols {
		allChars += specialChars

		mandatorySpecialChar := selectRandomCharacter(specialChars)
		password = append(password, mandatorySpecialChar)

	}

	length -= len(password)

	for i := 0; i < length; i++ {
		// Generate a random index to select a character from allChars
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}

		// Append the randomly selected character to the password
		password = append(password, string(allChars[randomIndex.Int64()]))
	}

	password = shuffle(password)
	passwordString := strings.Join(password, "")

	return passwordString, nil
}

func Generate(generateFlags *flag.FlagSet) (string, string) {
	generateFlags.BoolVar(&noSymbols, "n", false, "Skip symbols while generating password")
	generateFlags.BoolVar(&copyToClipBoard, "c", false, "Copy to clipboard.")
	generateFlags.Parse(os.Args[2:])
	passwordLength := 25
	genArgs := generateFlags.Args()

	if len(genArgs) > 1 {
		length, err := strconv.Atoi(genArgs[1])
		if err != nil {
			log.Println("Error: Password length should be an integer.")
			os.Exit(0)
		}
		passwordLength = length
	} else if len(genArgs) == 0 {
		log.Fatal("Usage: dost generate [-n] [-c] passwordName [passwordLength]")
	}

	passwordName := genArgs[0]

	// Generate and print the password
	password, err1 := generatePassword(passwordLength, noSymbols)
	if err1 != nil {
		log.Println("Error generating password:", err1)
	}

	if err1 == nil {
		if copyToClipBoard {
			err2 := writeToClipboard(password)
			if err2 != nil {
				log.Println("Error writing to clipboard:", err2)
				os.Exit(1)
			}
			fmt.Println("Copied to clipboard! ✅")
		} else {
			fmt.Println("Generated Password:", password)
		}
	}

	return password, passwordName

}
