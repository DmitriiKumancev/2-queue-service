.PHONY: build run

build:
	@go build -o 2queue-service

run:
	@go run .

test:
	@go test -v

clean:
	@rm -f 2queue-service


