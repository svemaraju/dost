package internal

import (
	"crypto/rand"
	"math/big"
)


// GeneratePassword returns a random password of the specified length
func GeneratePassword(length int) (string, error) {
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