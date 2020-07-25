VERSION=$(shell git describe --always --dirty --tags)

pushnix:
	go build -ldflags "-s -w -X main.Version=$(VERSION)" -o pushnix

clean:
	rm -f pushnix
