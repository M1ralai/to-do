package server

import (
	"log"
	"net/http"
)

type Server struct {
	serverAddr string
	mux        *http.ServeMux
}

func NewServer(addr string) *Server {
	return &Server{
		serverAddr: addr,
		mux:        nil,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	s.mux = mux
}

func (s *Server) Run() {
	s.mux.HandleFunc("/index", s.getIndex)
	http.ListenAndServe(s.serverAddr, s.mux)
}

func (s *Server) getIndex(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println(r.Method)
		log.Println(r.URL)
	case "POST":
		log.Println(r.Method)
		log.Println(r.URL)
	case "DELETE":
		log.Println(r.Method)
		log.Println(r.URL)
	case "PUT":
		log.Println(r.Method)
		log.Println(r.URL)
	default:
		log.Println(r.URL)
		log.Println("unauthorized url")
	}
}
