package main

import "net/http"

type StatusResponse struct {
	Message string `json:"message"`
}

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, StatusResponse{Message: "Server is running"})
}
