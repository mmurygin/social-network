include .env

build:
	packer build infra/build.json

# local development
local-env:
	. .local-env

local-init:
	go get ./...

local-run: local-env
	go run main.go

gin: local-env
	gin -b tmp/ginbin run main.go
