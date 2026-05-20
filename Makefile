include ./.env

all: run

run:
	go run ./src 

build:
	GOARCH=amd64 GOOS=windows go build -o bin/$(BINARY_NAME)-win.exe ./src
	GOARCH=amd64 GOOS=linux go build -o bin/$(BINARY_NAME)-lnx.exe ./src

clean:
	rm -rf bin/*.exe
