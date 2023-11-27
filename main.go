package main

import (
	"os"
    "crypto/rand"
    "fmt"
    "math/big"
    "github.com/atotto/clipboard"
    "github.com/svemaraju/dost/internal"
)


// writeToClipboard writes a given text to clipboard
func writeToClipboard(text string) error {
    return clipboard.WriteAll(text)
}


// generatePassword generates a random password of the specified length
func generatePassword(length int) (string, error) {
    const (
        uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
        digits           = "0123456789"
        specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
    )

    // Combine all characters for password generation
    allChars := uppercaseLetters + lowercaseLetters + digits + specialChars

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


func main() {
    args := argutils.ReadArgs()

	// Generate and print the password
	password, err1 := generatePassword(args.PasswordLength)
	if err1 != nil {
		fmt.Println("Error generating password:", err1)
		return
	}

	fmt.Println("Generated Password:", password)

	err2 := writeToClipboard(password)
	if err2 != nil {
		fmt.Println("Error writing to clipboard:", err2)
		os.Exit(1)
	}

	fmt.Println("Copied to clipboard! ✅\n")
}