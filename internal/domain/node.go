package domain

import (
	"errors"
)

var (
	ErrInvalidNode = errors.New("invalid node")
)

type Node struct {
	Addr string
	FileName string
}

func (node *Node) IsValid() (ok bool) {
	return true
}
