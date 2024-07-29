package internal

import (
	"crypto/rand"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
)

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

func Generate(generateFlags *flag.FlagSet, noSymbols bool) (string, error) {
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
		return password, err1
	}

	return password, nil
}
