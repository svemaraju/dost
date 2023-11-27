package main

import (
	"os"
    "fmt"
    "github.com/svemaraju/dost/internal"
)


func main() {
    args := internal.ReadArgs()

	// Generate and print the password
	password, err1 := internal.GeneratePassword(args.PasswordLength)
	if err1 != nil {
		fmt.Println("Error generating password:", err1)
		return
	}

	fmt.Println("Generated Password:", password)

	err2 := internal.WriteToClipboard(password)
	if err2 != nil {
		fmt.Println("Error writing to clipboard:", err2)
		os.Exit(1)
	}

	fmt.Println("Copied to clipboard! ✅")
}