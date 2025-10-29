package hash

import "github.com/multiformats/go-multihash"

// IsValidCode checks to see if the given multihash code is valid.
func IsValidCode(code uint64) bool {
	_, ok := multihash.Codes[code]
	return ok
}
