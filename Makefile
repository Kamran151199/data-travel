test:
	go test ./...
run:
	go run main.go
build:
	go build -o ./bin/dbmigrate main.go