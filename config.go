package main

import (
	"encoding/json"
	"os"
)

// Something like:
//
//	{
//	  "HOST": "http://192.168.1.0:8123",
//	  "TOKEN": "deadbeef",
//	  "SENSOR_NAMES": ["plant_sensor_7897"],
//	  "POLL_INTERVAL_MS": 5000,
//	}
//
// Create/retrieve your token from http://$HOST/profile.
type Config struct {
	Host           string   `json:"HOST"`
	Token          string   `json:"TOKEN"`
	SensorNames    []string `json:"SENSOR_NAMES"`
	PollIntervalMs int      `json:"POLL_INTERVAL_MS"`
}

func readConfig() Config {
	var bytes, err = os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	json.Unmarshal(bytes, &config)
	if len(config.Host) == 0 {
		panic("You must provide a \"HOST\" field in your `config.json` file.")
	}
	if len(config.Token) == 0 {
		panic("You must provide a \"TOKEN\" field in your `config.json` file.")
	}
	if config.PollIntervalMs == 0 {
		panic("You must provide a \"PollIntervalMs\" field in your `config.json` file.")
	}
	return config
}
