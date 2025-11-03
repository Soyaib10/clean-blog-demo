package comment

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/", h.Create).Methods("POST")
	r.HandleFunc("/post/{postID}", h.ListByPost).Methods("GET")
}
