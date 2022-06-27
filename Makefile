.PHONY: test vtest deps tidy

test: deps
	go test -race -count 1 ./...

vtest: deps
	go test -v -race -count 1 ./...

deps:
	go mod download

tidy:
	go mod tidy
