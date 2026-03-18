package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"transfile/internal/domain"
)

/*
hashes using sha256. O(n)
*/
func NewHash(input io.Reader, length int) (domain.Hash) {
	var s []byte
	for length > 0 {
		p := make([]byte, 1024)
		n, _ := input.Read(p)
		s = append(s, p[0:n]...)
		length -= n
	}
	h := sha256.Sum256(s)
	return domain.Hash(hex.EncodeToString(h[:]))
}
