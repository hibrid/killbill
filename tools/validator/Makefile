.PHONY: all check-deps install-deps-debian install-deps-macos

all: check-validator-deps

check-validator-deps:
ifeq ($(shell uname), Linux)
	@dpkg -l | grep libxml2-dev || make install-deps-debian
else ifeq ($(shell uname), Darwin)
	@brew list libxml2 || make install-deps-macos
else
	@echo "Unsupported OS"
endif

install-validator-deps-debian:
	@sudo apt-get update && sudo apt-get install -y build-essential libxml2 libxml2-dev

install-validator-deps-macos:
	@brew install libxml2

run:
	@go run main.go
