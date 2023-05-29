package albero

import (
	"crypto/sha256"
)

// Defines the chosen hasher function (up to developer to redefine this)
var hasherFunction = sha256.Sum256 // Hash function agnostic fashion

// Proxies the invocation of the hasher
func PerformHash(data []byte) []byte {
	h := hasherFunction(data)
	return h[:]
}
