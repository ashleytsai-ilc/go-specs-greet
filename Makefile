build:
	golangci-lint run
	go test ./...

unit-tests:
	go test -short ./...