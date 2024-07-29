package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/svemaraju/dost/internal"
)

var (
	noSymbols       bool
	copyToClipBoard bool = true
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateCmd.BoolVar(&noSymbols, "n", false, "Skip symbols while generating password")
	generateCmd.BoolVar(&copyToClipBoard, "c", false, "Copy to clipboard.")
	flag.Parse()

	switch os.Args[1] {
	case "generate":
		password, err := internal.Generate(generateCmd, noSymbols)
		if err == nil {
			if copyToClipBoard {
				err2 := internal.WriteToClipboard(password)
				if err2 != nil {
					fmt.Println("Error writing to clipboard:", err2)
					os.Exit(1)
				}
				fmt.Println("Copied to clipboard! ✅")
			} else {
				fmt.Println("Generated Password:", password)
			}
		}
	}
}
