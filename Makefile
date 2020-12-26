# Load any variables in .env and export them
include .env
export

BIN := bin
APP := $(BIN)/godeletemytweets

build:
	@go build -o $(APP) main.go

run: build
	@$(APP)
