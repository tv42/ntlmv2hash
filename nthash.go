package ntlmv2hash

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/md4"
	"unicode/utf16"
)

// NTPasswordHash computes the NTLM v2 password hash.
//
// The output is password-equivalent and easy to reverse. It must be
// guarded just as well as the original password.
func NTPasswordHash(password string) string {
	input := utf16.Encode([]rune(password))
	h := md4.New()
	if err := binary.Write(h, binary.LittleEndian, input); err != nil {
		// these are all in-memory operations with no error modes,
		// but just in case
		panic(fmt.Errorf("impossible error hashing password: %w", err))
	}
	output := h.Sum(nil)
	// encode to conventional uppercase hex
	return fmt.Sprintf("%X", output)
}
