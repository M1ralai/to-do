build: compile run

compile:
	go build -o bin/out cmd/main.go

run:
	./bin/out
