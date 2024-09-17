package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/svemaraju/dost/internal"
)

func printHelp() {
	fmt.Println("Invalid command to run dost password manager.")
	fmt.Println("Please choose one of the following options:")
	fmt.Println("dost init")
	fmt.Println("dost generate [-c] [-n] <password-name>")
	fmt.Println("dost show <password-name>")
}

func main() {
	path := os.Getenv("HOME") + "/.dost"
	storage := internal.GetStorage(path)
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	flag.Parse()

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		if storage.IsReady() {
			password, passwordName := internal.Generate(generateCmd)
			if err := storage.Add(password, passwordName); err != nil {
				fmt.Printf("%v", err)
			}
		} else {
			fmt.Println("dost is not initialized. Run `dost init`")
			os.Exit(1)
		}
	case "init":
		storage.Init()
	case "show":
		password, err := storage.Show(os.Args[2])
		if err != nil {
			fmt.Printf("Something went wrong: %v", err)
		} else {
			fmt.Println(password)
		}
	default:
		printHelp()

	}
}
