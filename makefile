.PHONY: default run build test docs clean

# Variables
APP_NAME=github.com/thiago-s-silva/ms-go-task-management-api

# Tasks
default: run

run:
	@go run cmd/main.go