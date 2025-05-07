package main

import (
	"fmt"
	"io"
	"net/http"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func main() {
	var config = readConfig()
	var endpoint = fmt.Sprintf("%s/api/states", config.Host)
	initRender()
	defer endRender()
	for !rl.WindowShouldClose() {
		var response = request(config, endpoint)
		var sensorStats = UnmarshallStates(config, response)
		render(sensorStats)
		// TODO move to another goroutine
		time.Sleep(time.Millisecond * time.Duration(config.PollIntervalMs))
	}
}

func request(config Config, url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 300 {
		var msg = fmt.Sprintf("Error code %d\nAttempting \"%s\"\n%s", resp.StatusCode, url, body)
		panic(msg)
	}
	return body
}
