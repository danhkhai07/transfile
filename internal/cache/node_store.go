package cache

import (
	"hash"
	"io"
	"transfile/internal/domain"
	"transfile/internal/util"
)

type FileStore struct {
	nodeMapSize int
	nodeMap map[domain.Hash][]domain.Node
	sizeMap map[domain.Hash]int64
}

func NewFileStore() (*FileStore) {
	store := FileStore{
		nodeMap: make(map[domain.Hash][]domain.Node),
		sizeMap: make(map[domain.Hash]int64),
	}
	return &store
}

func (mcs *FileStore) GetNodes(hash domain.Hash) (nodes []domain.Node, ok bool) {
	nodes, ok = mcs.nodeMap[hash]
	return nodes, ok
}

func (mcs *FileStore) GetNumberOfNodes(hash domain.Hash) (size int, ok bool) {
	nodes, ok := mcs.nodeMap[hash]
	return len(nodes), ok
}

func (mcs *FileStore) GetFileSize(hash domain.Hash) (size int64, ok bool) {
	size, ok = mcs.sizeMap[hash]
	return size, ok
}

func (mcs *FileStore) addNode(hash domain.Hash, addr string, filename string) (err error) {
	node := domain.Node{
		Addr: addr,
		FileName: filename,
	}
	if !node.IsValid() {
		return domain.ErrInvalidNode
	}

	mcs.nodeMap[hash] = append(mcs.nodeMap[hash], node)
	mcs.nodeMapSize += 1
	return nil
}

// 		hash: "abc123",
// 		node_addr: "192.168.1.1:52000"
// 		file_name: "Never_Gonna_Give_U_Up.mp4"
// 		size: 734003200
func (mcs *FileStore) AddFile(
	hash domain.Hash,
	nodeAddr string,
	fileName string,
	size int64,
) (err error) {
	if !hash.IsValid() {
		return domain.ErrInvalidHash
	}

	node := domain.Node{
		Addr: nodeAddr,
		FileName: fileName,
	}
	if !node.IsValid() {
		return domain.ErrInvalidNode
	}

	mcs.sizeMap[hash] = size
	
}
