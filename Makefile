all:
	go build -o ./build/goserver ./cmd/goserver/server.go
	go build -o ./build/goclient ./cmd/goclient/client.go
