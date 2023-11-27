package internal

import (
	"flag"
)


type argset struct {
    PasswordLength int
}


func ReadArgs() argset {
    // Read the desired password length
    passwordLength := flag.Int("length", 12, "length of your password")
    flag.Parse()

    return argset{
        PasswordLength: *passwordLength,
    }

}