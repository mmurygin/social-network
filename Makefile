include .env

build:
	packer build infra/build.json

# local development
init:
	go get ./...

run:
	go run main.go

gin:
	gin -b tmp/ginbin run main.go
