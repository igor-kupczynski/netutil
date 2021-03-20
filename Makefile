all: netutil-serve

.PHONY: netutil-serve
netutil-serve:
	go build -o . -v ./...
