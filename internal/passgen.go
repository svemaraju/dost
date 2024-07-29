package internal

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

var (
	noSymbols       bool
	copyToClipBoard bool = true
)

func writeToClipboard(text string) error {
	return clipboard.WriteAll(text)
}

// generatePassword returns a random password of the specified length
func generatePassword(length int, noSymbols bool) (string, error) {
	const (
		uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
		digits           = "0123456789"
		specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
	)

	// Combine all characters for password generation
	allChars := uppercaseLetters + lowercaseLetters + digits

	if !noSymbols {
		allChars += specialChars
	}

	var password string
	for i := 0; i < length; i++ {
		// Generate a random index to select a character from allChars
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}

		// Append the randomly selected character to the password
		password += string(allChars[randomIndex.Int64()])
	}

	return password, nil
}

func Generate(generateFlags *flag.FlagSet) {
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
	}

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

}
