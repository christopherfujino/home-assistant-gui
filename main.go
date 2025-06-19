package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"io"
	"net/http"
	"os" // for os.Getenv
	"strconv"
	"strings"
	"time"
)

func main() {
	var config = readConfig()
	var endpoint = fmt.Sprintf("%s/api/states", config.Host)
	var statsChan = make(chan Stats)

	go poll(config, endpoint, statsChan)

	initRender(config)
	defer endRender()
	var currentStats Stats
	for !rl.WindowShouldClose() {
		select {
		case currentStats = <-statsChan:
			fmt.Printf("Received new stats from channel\n")
			// no-op
		}
		render(currentStats)
	}
}

func poll(config Config, url string, sensorChan chan Stats) {
	// /sys/class/thermal/thermal_zone0/temp
	if os.Getenv("MOCK_API") == "1" {
		fmt.Printf("Sent mock data\n")
		sensorChan <- Stats{
			SensorStats: SensorStats{
				Battery:     69,
				Temperature: 69.69,
				Illuminance: 420,
			},
			CPUTemp: 42,
		}
	} else {
		for {
			var response = request(config, url)
			var sensorStats = UnmarshallStates(config, response)
			sensorChan <- Stats{
				SensorStats: sensorStats,
				CPUTemp:     readTemp(),
			}
			time.Sleep(time.Millisecond * time.Duration(config.PollIntervalMs))
		}
	}
}

func readTemp() float64 {
	bytes, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}
	s := string(bytes)
	s = strings.TrimSpace(s)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	f = f / 1000
	return f
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
