.DEFAULT_GOAL := build

all : test doc clean build
.PHONY: all

build:
	go build -ldflags="-X 'main.Version=v0.0.3'" -o propgen main.go parsing.go

test:
	go test ./pkg

doc:
	godoc -http=:6060

clean:
	rm propgen 2> /dev/null || true
