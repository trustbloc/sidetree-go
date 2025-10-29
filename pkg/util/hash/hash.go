package hash

import "github.com/multiformats/go-multihash"

// IsValidCode checks to see if the given multihash code is valid.
func IsValidCode(code uint64) bool {
	switch code {
	case multihash.SHA2_256, multihash.SHA2_512:
		return true
	default:
		return false
	}
}
