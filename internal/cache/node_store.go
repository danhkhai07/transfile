package cache

import (
	"transfile/internal/domain"
)

type NodeStore struct {
	size int
	fileMap map[domain.Hash][]domain.Node
}

func NewNodeStore() (*NodeStore) {
	store := NodeStore{
		fileMap: make(map[domain.Hash][]domain.Node),
	}
	return &store
}

func (mcs *NodeStore) GetNodes(hash domain.Hash) (nodes []domain.Node, ok bool) {
	nodes, ok = mcs.fileMap[hash]
	return nodes, ok
}

func (mcs *NodeStore) AddNode(hash domain.Hash, addr string) (err error) {
	node := domain.Node{
		Addr: addr,
	}
	if !node.IsValid() {
		return domain.ErrInvalidNode
	}

	mcs.fileMap[hash] = append(mcs.fileMap[hash], node)
	mcs.size += 1
	return nil
}
