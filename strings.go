package gorand

import (
	"crypto/rand"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
)

// GetRandomChars returns a fixed length string of chars contained on a collection string provided.
// Params:
// - c (string): Collection of chars to randomly pick from.
// - n (int): Length of the result string
func GetRandomChars(c string, n int) (string, error) {
	var r string

	for i := 0; i < n; i++ {
		// Read random position
		p, err := rand.Int(rand.Reader, big.NewInt(int64(len(c))))
		if err != nil {
			return "", err
		}

		r += string(c[p.Int64()])
	}

	return r, nil
}

// GetAlphaNumString returns a fixed length string of random letters and numbers [a-z][A-Z][0-9]
// Params:
// - n (int): Length of the result string
func GetAlphaNumString(n int) (string, error) {
	return GetRandomChars(lowercase+uppercase+numbers, n)
}

// GetAlphaString returns a fixed length string of random letters [a-z][A-Z]
// Params:
// - n (int): Length of the result string
func GetAlphaString(n int) (string, error) {
	return GetRandomChars(lowercase+uppercase, n)
}
