package main

import "github.com/M1ralai/todo/cmd/server"

func main() {
	s := server.NewServer(":3000")
	s.Start()
	s.Run()
}
