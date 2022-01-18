// Package random provides utilities for creating random things
package random

import (
	cryptorand "crypto/rand"
	"math/big"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// WeakRandString - returns a random string of the given length
func WeakRandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))] //nolint:gosec //reason: don't need strong crypto here
	}
	return string(b)
}

// SecureRandStringFromChars - returns a crypto safe random string of the given length with the given chars
func SecureRandStringFromChars(n int, letters string) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// SecureRandString - returns a crypto safe random string of the given length with the default chars
func SecureRandString(n int) (string, error) {
	return SecureRandStringFromChars(n, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-")
}
