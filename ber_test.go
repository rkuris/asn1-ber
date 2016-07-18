package ber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	b := EncodeInteger(0)
	assert.Equal(t, []byte{0}, b, "Encode of 0 failed")

	b = EncodeInteger(127)
	assert.Equal(t, []byte{0x7f}, b, "Encode of 127 failed")

	b = EncodeInteger(128)
	assert.Equal(t, []byte{0, 0x80}, b, "Encode of 128 failed")

	b = EncodeInteger(256)
	assert.Equal(t, []byte{1, 0}, b, "Encode of 256 failed")
}
