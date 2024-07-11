package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
