package server

import (
	"transfile/internal/domain"
)

type getHealthRequest struct {}

type getHealthResponse struct {
	Status string `json:"status"`
	Uptime int `json:"uptime"`
}

type lookupFileRequest struct {}

type lookupFileResponse struct {
	Found bool `json:"found"`
	Hash domain.Hash `json:"file_hash"`
	Size int64 `json:"size"`
	NumberOfNodes int `json:"number_of_nodes"`
	Nodes []domain.Node `json:"nodes"`
}

type downloadFileRequest struct {
	
}

type downloadFileResponse struct {
	
}

// 		file_hash: "abc123",
// 		node_addr: "192.168.1.1:52000"
// 		file_name: "Never_Gonna_Give_U_Up.mp4"
// 		size: 734003200
type postFileRequest struct {
	Hash domain.Hash `json:"file_hash"`
	NodeAddr string `json:"node_addr"`
	FileName string `json:"file_name"`
	Size int64 `json:"size"`
}

type postFileResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
