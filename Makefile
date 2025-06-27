.PHONY: all build clean bin

bin/silo: bin tidy
	go build -o bin/silo ./cmd/main.go

test:
	go test -v ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/silo

bin:
	mkdir -p bin



