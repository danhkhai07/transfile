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

type postFileRequest struct {
	
}

type postFileResponse struct {
	
}
