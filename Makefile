.PHONY: run
run:
	LIBGL_ALWAYS_SOFTWARE=1 go run . 2>&1 | tee ./debug.log

# This was only for the video
.PHONY: json
json:
	go run . -- json


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
