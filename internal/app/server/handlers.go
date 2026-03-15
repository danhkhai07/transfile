package server

import (
	"encoding/json"
	"net/http"
)

// GET /health
func (svr *Server) getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := getHealthResponse{
		Status: "ok",
		Uptime: svr.Uptime(),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		svr.logger.Errwriteln("json parsing error: %s", err)
	}
}

// GET /lookup/{hash}
func (svr *Server) lookupFile(w http.ResponseWriter, r *http.Request) {

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
