.PHONY: test

test:
	go test ./... -v -short

test-long:
	go test ./...