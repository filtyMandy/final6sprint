package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
)

type Server struct {
	logger *log.Logger
	srv    *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	s := &http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: logger,
	}

	return &Server{
		logger: logger,
		srv:    s,
	}
}

func (s *Server) Start() error {
	s.logger.Println("Starting server on", s.srv.Addr)
	return s.srv.ListenAndServe()
}
