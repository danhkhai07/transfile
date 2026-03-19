package server

import (
	"encoding/json"
	"net/http"
	"transfile/internal/domain"
)

// GET /health
func (svr *Server) getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := getHealthResponse{
		Status: "ok",
		Uptime: svr.Uptime(),
	}

	if err := svr.Encode(w, r, http.StatusOK, resp); err != nil {
		svr.logger.Errwriteln("encoding json: %s", err)
	}
}

// GET /lookup/{hash}
func (svr *Server) lookupFile(w http.ResponseWriter, r *http.Request) {
	hash := domain.Hash(r.PathValue("hash"))
	
	resp := lookupFileResponse{
		Found: false,
		Hash: hash,
		Size: -1,
		NumberOfNodes: -1,
		Nodes: nil,
	}
	nodes, ok := svr.fileStore.GetNodes(hash)
	if !ok {
		if err := svr.Encode(w, r, http.StatusOK, resp); err != nil {
			svr.logger.Errwriteln("encode json: %s", err)
		}
		return
	}
	
	fileSize, ok := svr.fileStore.GetFileSize(hash)
	if !ok {
		fileSize = -1
	}
	numsOfNodes, ok := svr.fileStore.GetNumberOfNodes(hash)
	if !ok {
		numsOfNodes = -1
	}

	resp = lookupFileResponse{
		Found: true,
		Hash: hash,
		Size: fileSize,
		NumberOfNodes: numsOfNodes,
		Nodes: nodes,
	}

	if err := svr.Encode(w, r, http.StatusOK, resp); err != nil {
		svr.logger.Errwriteln("encode json: %s", err)
	}
}

// GET /download/{hash}
func (svr *Server) downloadFile(w http.ResponseWriter, r *http.Request) {

}

// POST /upload
// {
// 		hash: "abc123",
// 		node_addr: "192.168.1.1:52000"
// 		file_name: "Never_Gonna_Give_U_Up.mp4"
// 		size: 734003200
// }
func (svr *Server) postFile(w http.ResponseWriter, r *http.Request) {
	
}
