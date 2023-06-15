package albero

import (
	"crypto/sha256"
)

// Defines the chosen hasher function (up to developer to redefine this)
var hasherFunction = sha256.Sum256 // Hash function agnostic fashion

/*
- Description: Computes the hash value for the provided input data using the chosen hash function.
- Input: data - the input data to be hashed.
- Output: Returns the computed hash value as a byte slice.
*/
func PerformHash(data []byte) []byte {
	h := hasherFunction(data)
	return h[:]
}
