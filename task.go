package main

import (
	"math/rand/v2"
	"strings"
	"uuid"
)

const (
	upperLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters string = "abcdefghijklmnopqrstuvwxyz"
	digits       string = "0123456789"
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
