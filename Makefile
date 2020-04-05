include .env

init:
	go get ./...

run:
	go run main.go

gin:
	gin -b tmp/ginbin run main.go
