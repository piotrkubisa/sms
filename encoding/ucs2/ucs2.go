package ucs2

import (
	"encoding/binary"
	"errors"
	"fmt"
	"unicode/utf16"
)

// Decode converts an array of UCS2 characters into an array of runes.
// As the UCS2 characters are packed into a byte array, the length of the
// byte array provided must be even.
func Decode(src []byte) ([]rune, error) {
	if len(src) == 0 {
		return nil, nil
	}
	if len(src)&0x01 == 0x01 {
		return nil, ErrInvalidLength
	}
	l := len(src) / 2
	dst := make([]rune, 0, l)
	for ri := 0; ri < len(src)-1; ri = ri + 2 {
		r := rune(binary.BigEndian.Uint16(src[ri:]))
		if utf16.IsSurrogate(r) {
			ri += 2
			if ri >= len(src)-1 {
				return dst, ErrDanglingSurrogate(uint16(r))
			}
			r2 := rune(binary.BigEndian.Uint16(src[ri:]))
			r = utf16.DecodeRune(r, r2)
		}
		dst = append(dst, r)
	}
	return dst, nil
}

// Encode converts an array of UCS2 runes into an array of bytes, where pairs of
// bytes (in Big Endian) represent a UCS2 character.
func Encode(src []rune) []byte {
	if len(src) == 0 {
		return nil
	}
	u := utf16.Encode(src)
	dst := make([]byte, len(u)*2)
	wi := 0
	for _, r := range u {
		binary.BigEndian.PutUint16(dst[wi:], uint16(r))
		wi += 2
	}
	return dst
}

// ErrDanglingSurrogate indicates only half of a suggorate pair is provided at
// the end of the byte array being decoded.
type ErrDanglingSurrogate uint16

func (e ErrDanglingSurrogate) Error() string {
	return fmt.Sprintf("ucs2: dangling surrogate: 0x%04x", uint16(e))
}

// Surrogate returns the dangling surrogate.
func (e ErrDanglingSurrogate) Surrogate() uint16 {
	return uint16(e)
}

var (
	// ErrInvalidLength indicates the binary provided has an invalid (odd) length.
	ErrInvalidLength = errors.New("ucs2: length must be even")
)
