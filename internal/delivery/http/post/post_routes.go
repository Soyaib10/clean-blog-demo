package post

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("", h.Create).Methods("POST")
	r.HandleFunc("/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/user/{userID}", h.ListByUser).Methods("GET")
}
