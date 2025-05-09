package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"io"
	"net/http"
	"time"
)

func main() {
	var config = readConfig()
	var endpoint = fmt.Sprintf("%s/api/states", config.Host)
	var statsChan = make(chan SensorStats)
	go poll(config, endpoint, statsChan)
	initRender(config)
	defer endRender()
	var currentStats SensorStats
	for !rl.WindowShouldClose() {
		select {
		case currentStats = <-statsChan:
			fmt.Printf("Received new stats from channel\n")
			// no-op
		default:
			// Do not block
		}
		render(currentStats)
	}
}

func poll(config Config, url string, sensorChan chan SensorStats) {
	for {
		var response = request(config, url)
		var sensorStats = UnmarshallStates(config, response)
		sensorChan <- sensorStats
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
		// TODO surface to UI that there was a network error
		var msg = err.Error()
		panic(msg)
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
