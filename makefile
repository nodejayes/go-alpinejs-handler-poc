build:
	templ generate && go build -o bin/server.exe cmd/main.go
run:
	./bin/server.exe