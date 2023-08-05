package main

import (
	"crypto/rand"
	"math/big"
)

const (
	symbols        = "!@#$%^&*"
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrstuvwxyz"
	numbers        = "0123456789"
)

func generatePassword(length int) (string, error) {
	chars := capitalLetters + smallLetters + numbers + symbols

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
