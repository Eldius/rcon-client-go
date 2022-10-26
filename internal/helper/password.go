package helper

import (
	"fmt"
	"golang.org/x/term"
	"strings"
	"syscall"
)

func AskForPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}
