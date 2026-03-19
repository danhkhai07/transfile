package domain

import (
	"io"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

var (
	ErrInvalidHash = errors.New("invalid hash")
)

type Hash string

func NewHash(input io.Reader, length int) (Hash) {
	var s []byte
	for length > 0 {
		p := make([]byte, 1024)
		n, _ := input.Read(p)
		s = append(s, p[0:n]...)
		length -= n
	}
	hsh := sha256.Sum256(s)
	return Hash(hex.EncodeToString(hsh[:]))
}

func (h *Hash) IsValid() (ok bool) {
	if len(*h) != 64 {
		return false
	}
	for _, c := range *h {
		if !((c >= '0' && c <= '9') ||
			(c >= 'a' && c <= 'f') ||
			(c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func (h *Hash) Equals(input io.Reader, length int) (ok bool) {
	var s []byte
	for length > 0 {
		p := make([]byte, 1024)
		n, _ := input.Read(p)
		s = append(s, p[0:n]...)
		length -= n
	}
	hsh := sha256.Sum256(s)
	return hex.EncodeToString(hsh[:]) == string(*h)
}
