package server

import (
	
)

type getHealthRequest struct {}

type getHealthResponse struct {
	Status string `json:"status"`
	Uptime int `json:"uptime"`
}

type lookupFileRequest struct {
	
}

type lookupFileResponse struct {

}

type downloadFileRequest struct {
	
}

type downloadFileResponse struct {
	
}

type postFileRequest struct {
	
}

type postFileResponse struct {
	
}
