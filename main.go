package main

import (
	"flag"
	"os"

	"github.com/svemaraju/dost/internal"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	flag.Parse()

	switch os.Args[1] {
	case "generate":
		internal.Generate(generateCmd)

	}
}
