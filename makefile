.PHONY: run build clean

run:
	go run main.go

build:
	go build -o telegram-bot main.go

clean:
	rm -f telegram-bot

test:
	go test ./...

format:
	gofmt -w .
