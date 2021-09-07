.PHONY: client
client:
	@echo "Building the client bin"
	go build -o bin/client cmd/client/main.go