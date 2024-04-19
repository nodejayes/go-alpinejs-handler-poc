build:
	go build -o bin/server.exe cmd/main.go
test:
	go test ./...
test_context_store:
	go test ./context_store/...
run:
	./bin/server.exe