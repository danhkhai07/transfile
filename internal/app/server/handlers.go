package app

import "net/http"

// GET /health
func (svr *Server) getHealth(http.ResponseWriter, *http.Request) {

}

// GET /lookup/{hash}
func (svr *Server) lookupFile(http.ResponseWriter, *http.Request) {

}

// GET /download/{hash}
func (svr *Server) downloadFile(http.ResponseWriter, *http.Request) {

}

// POST /upload
// {
// 	hash: "abc123",
// 	node_addr: "192.168.1.1:52000"
// 	file_name: "Never_Gonna_Give_U_Up.mp4"
// 	size: 734003200
// 	
// }
func (svr *Server) postFile(http.ResponseWriter, *http.Request) {

}
