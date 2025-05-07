package main

import (
	"bytes"
	"encoding/json"
)

func format(raw []byte) string {
	const prefix = ""
	const indent = "  "
	var buffer bytes.Buffer
	json.Indent(&buffer, raw, prefix, indent)
	return buffer.String()
}
