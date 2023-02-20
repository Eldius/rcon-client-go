/*
Package helper holds some helper functions
*/
package helper

import (
	"fmt"
	"golang.org/x/term"
	"strings"
	"syscall"
)

// AskForPassword password ask helper function
func AskForPassword(prompt string) (string, error) {
	fmt.Printf("%s ", prompt)
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}
