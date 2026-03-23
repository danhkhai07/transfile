package server

import (
	"net/http"
	"transfile/internal/domain"
)

// GET /health
func (svr *Server) getHealth(w http.ResponseWriter, r *http.Request) {
	resp := getHealthResponse{
		Status: "ok",
		Uptime: svr.Uptime(),
	}

	if err := svr.Encode(w, r, http.StatusOK, resp); err != nil {
		svr.logger.Errwriteln("encode json: %s", err)
		return
	}
}

// GET /lookup/{hash}
func (svr *Server) lookupFile(w http.ResponseWriter, r *http.Request) {
	hash := domain.Hash(r.PathValue("hash"))
	
	resp := lookupFileResponse{
		Found: false,
		Hash: hash,
		Size: 0,
		NumberOfNodes: 0,
		Nodes: nil,
	}
	nodes, ok := svr.fileStore.GetNodes(hash)
	if !ok {
		if err := svr.Encode(w, r, http.StatusBadRequest, resp); err != nil {
			svr.logger.Errwriteln("encode json: %s", err)
		}
		return
	}
	
	fileSize, ok := svr.fileStore.GetFileSize(hash)
	if !ok {
		fileSize = 0
	}
	numsOfNodes, ok := svr.fileStore.GetNumberOfNodes(hash)
	if !ok {
		numsOfNodes = 0
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
		return
	}
}

// GET /download/{hash}
func (svr *Server) downloadFile(w http.ResponseWriter, r *http.Request) {

}

// POST /upload
// {
// 		file_hash: "abc123",
// 		node_addr: "192.168.1.1:52000"
// 		file_name: "Never_Gonna_Give_U_Up.mp4"
// 		size: 734003200
// }
func (svr *Server) postFile(w http.ResponseWriter, r *http.Request) {
	req := postFileRequest{}
	if err := svr.Decode(r, &req); err != nil {
		svr.logger.Errwriteln("decode json: %s", err)
		return
	}

	resp := postFileResponse{
		Status: "success",
		Message: "",
	}
	if req.Hash == "" || req.NodeAddr == "" || req.Size <= 0 {
		resp.Status = "failed"
		resp.Message = "missing fields"
		if err := svr.Encode(w, r, http.StatusBadRequest, resp); err != nil {
			svr.logger.Errwriteln("encode json: %s", err)
		}
		return
	}
	if req.FileName == "" {
		req.FileName = "Unnamed_file"
	}

	err := svr.fileStore.AddFile(req.Hash, req.NodeAddr, req.FileName, req.Size)
	if err != nil {
		resp.Status = "failed"
		resp.Message = err.Error()
		if err := svr.Encode(w, r, http.StatusBadRequest, resp); err != nil {
			svr.logger.Errwriteln("encode json: %s", err)
		}
		return
	}

	if err := svr.Encode(w, r, http.StatusOK, resp); err != nil {
		svr.logger.Errwriteln("encode json: %s", err)
		return
	}
}
