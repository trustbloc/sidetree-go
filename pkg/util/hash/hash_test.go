package hash

import (
	"testing"

	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func TestIsValidCode(t *testing.T) {
	t.Run("code is valid SHA2_256", func(t *testing.T) {
		code := uint64(multihash.SHA2_256)
		require.True(t, IsValidCode(code))
	})
	t.Run("code is valid SHA2_512", func(t *testing.T) {
		code := uint64(multihash.SHA2_512)
		require.True(t, IsValidCode(code))
	})

	t.Run("code is not valid", func(t *testing.T) {
		code := uint64(123)
		require.False(t, IsValidCode(code))
	})
}
