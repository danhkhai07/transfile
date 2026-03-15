package app

import "net/http"

// GET /health
func getHealth() func(http.ResponseWriter, *http.Request) {
	return nil
}

// GET /lookup/{hash}
func lookupFile() func(http.ResponseWriter, *http.Request) {
	return nil
}

// GET /download/{hash}
func downloadFile() func(http.ResponseWriter, *http.Request) {
	return nil
}

// POST /upload
// {
// 	hash: "abc123",
// 	node_addr: "192.168.1.1:52000"
// 	file_name: "Never_Gonna_Give_U_Up.mp4"
// 	size: 734003200
// 	
// }
func postFile() func(http.ResponseWriter, *http.Request) {
	return nil
}
