package main

import (
	"github.com/M1ralai/todo/cmd/db"
	"github.com/M1ralai/todo/cmd/server"
)

func main() {
	db.InitDatabase()
	s := server.NewServer(":3000")
	s.Start()
	s.Run()
}
