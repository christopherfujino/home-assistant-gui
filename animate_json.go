package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"strings"
)

func animateJson() {
	var config = readConfig()
	var endpoint = fmt.Sprintf("%s/api/states", config.Host)

	var rawBytes = request(config, endpoint)

	var buffer = bytes.Buffer{}
	json.Indent(&buffer, rawBytes, "", "  ")
	animate(buffer.String())
}

func animate(src string) {
	const interval = time.Millisecond * 30
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		fmt.Printf("%s", line)
		time.Sleep(interval)
		fmt.Printf("\n")
		time.Sleep(interval)
	}
}
