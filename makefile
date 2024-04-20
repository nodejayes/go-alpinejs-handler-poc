build:
	go build -o bin/server.exe cmd/main.go
test:
	go test ./...
run:
	./bin/server.exe