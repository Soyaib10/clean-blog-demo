package user

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
    r.HandleFunc("", h.Create).Methods("POST")     
    r.HandleFunc("", h.List).Methods("GET")        
    r.HandleFunc("/{id}", h.GetByID).Methods("GET") 
}

