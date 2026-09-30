package handlers

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /user", GetUser)
	mux.HandleFunc("POST /user", CreateUser)
}
