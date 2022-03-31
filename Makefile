all:
	go build -o ./build/goserver ./pkg/server.go
	go build -o ./build/goclient ./pkg/client.go
