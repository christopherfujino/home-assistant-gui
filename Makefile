.PHONY: run
run:
	go run .

.PHONY: build
build:
	go build .

.PHONY: get
get: go.sum ;

go.sum: go.mod
	go get .

.PHONY: clean
clean:
	rm -f home-assistant-gui-go
