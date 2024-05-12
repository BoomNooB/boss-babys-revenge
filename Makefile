run:
	go mod tidy
	go run ./cmd/app/main.go 

test:
	go test ./... -cover -v 