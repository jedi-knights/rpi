
clean:
	rm -f tds

deps:
	go mod tidy
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build: clean
	go build -ldflags "-w -s" -o tds main.go

test:
	ginkgo ./...

lint:
	golangci-lint run ./...