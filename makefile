.PHONY: default run build test docs clean

# Variables
APP_NAME=github.com/thiago-s-silva/ms-go-task-management-api

# Tasks
default: run

run:
	@go run cmd/main.go

build:
	@CGO_ENABLED=0 GOOS=linux go build -o /main

up:
	@docker-compose up -d

down:
	@docker-compose up -d

clean-image:
	@docker rmi ms-go-task-management-api-api