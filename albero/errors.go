package albero

import "errors"

// Definitions of critical errors
var (
	ErrInvalidInput = errors.New("Invalid input")
	ErrInvalidProof = errors.New("Invalid proof")
	ErrInvalidTree  = errors.New("Invalid tree")
)
