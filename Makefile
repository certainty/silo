.PHONY: all build clean bin

bin/silo: bin
	go build -o bin/silo ./cmd/silo/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/silo

bin:
	mkdir -p bin



