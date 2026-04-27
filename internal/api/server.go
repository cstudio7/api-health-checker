package api

import (
	"encoding/json"
	"net/http"

	"api-health-checker/internal/repository"
)

type Server struct {
	repo repository.Repository
}

func NewServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) Start() {
	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(s.repo.GetAll())
	})

	http.ListenAndServe(":8080", nil)
}
