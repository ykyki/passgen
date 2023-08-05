package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	symbols        = "!@#$%^&*"
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrstuvwxyz"
	numbers        = "0123456789"
)

func generatePassword(length int, letterOption letterOption) (string, error) {
	var chars string
	if letterOption.capitalLetter {
		chars += capitalLetters
	}
	if letterOption.smallLetter {
		chars += smallLetters
	}
	if letterOption.number {
		chars += numbers
	}
	if letterOption.symbol {
		chars += symbols
	}

	if len(chars) == 0 {
		return "", errors.New("no letter option is specified")
	}

	var password string

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password += string(chars[randomIndex.Int64()])
	}

	return password, nil
}
