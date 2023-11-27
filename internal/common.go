package internal

import (
    "github.com/atotto/clipboard"
)


// WriteToClipboard adds a given text to clipboard
func WriteToClipboard(text string) error {
    return clipboard.WriteAll(text)
}