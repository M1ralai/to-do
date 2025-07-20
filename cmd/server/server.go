package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/M1ralai/todo/cmd/db"
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
	log.Println("server succesfully started")
}

func (s *Server) Run() {
	log.Println("server listening and serving at ", s.serverAddr)
	fs := http.FileServer(http.Dir("./html"))
	s.mux.Handle("/", http.StripPrefix("/", fs))
	s.mux.HandleFunc("/index", s.getIndex)
	http.ListenAndServe(s.serverAddr, s.mux)
}

func (s *Server) getIndex(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		data, err := db.GetTodo()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	case "POST":
		query := r.URL.Query()
		title := query.Get("title")
		desc := query.Get("desc")
		year := query.Get("year")
		month := query.Get("month")
		day := query.Get("day")
		hour := query.Get("hour")
		minute := query.Get("minute")
		date := year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + "00"
		layout := "2006-01-02 15:04:05" // Go'nun sabit layout formatı
		t, err := time.Parse(layout, date)
		if err != nil {
			fmt.Println("Hata:", err)
			return
		}
		db.CreateTodo(title, desc, t)

	case "DELETE":
		query := r.URL.Query()
		id := query.Get("id")
		cid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println(err)
		}
		db.DeleteTodo(int(cid))
	case "PUT":
		log.Println(r.Method)
		log.Println(r.URL)
	default:
		log.Println(r.URL)
		log.Println("unauthorized url")
	}
}
