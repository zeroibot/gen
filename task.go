package main

import (
	"math/rand/v2"
	"strings"
	"uuid"

	"golang.org/x/crypto/bcrypt"
)

const (
	upperLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters string = "abcdefghijklmnopqrstuvwxyz"
	digits       string = "0123456789"
	hashCost     int    = 10
)

func generateUUIDv7() string {
	return uuid.NewV7().String()
}

func generateCode(length int, pattern string) string {
	useUpper := strings.Contains(pattern, "A")
	useLower := strings.Contains(pattern, "a")
	useDigit := strings.Contains(pattern, "1")
	return generateRandomString(length, useUpper, useLower, useDigit)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func generateRandomString(length int, useUpper, useLower, useDigit bool) string {
	charSource := ""
	if useUpper {
		charSource += upperLetters
	}
	if useLower {
		charSource += lowerLetters
	}
	if useDigit {
		charSource += digits
	}
	numChars := len(charSource)
	if numChars == 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range length {
		idx := rand.IntN(numChars)
		b[i] = charSource[idx]
	}
	return string(b)
}
