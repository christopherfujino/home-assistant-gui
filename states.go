package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type SensorStats struct {
	Illuminance  int
	Battery      int
	Temperature  float64
	Moisture     int
	Conductivity int
}

type Stats struct {
	SensorStats
	CPUTemp float64
}

func UnmarshallStates(config Config, rawStates []byte) SensorStats {
	var states []map[string]any

	var illuminance *int
	var battery *int
	var temperature *float64
	var moisture *int
	var conductivity *int

	json.Unmarshal(rawStates, &states)
	for _, state := range states {
		var entityId = state["entity_id"].(string)
		for _, sensor := range config.SensorNames {
			var separator = fmt.Sprintf("sensor.%s_", sensor)
			if strings.HasPrefix(entityId, separator) {
				var _, after, found = strings.Cut(entityId, separator)
				if !found {
					panic("Huh?!")
				}

				switch after {
				case "moisture":
					moistureValue, err := strconv.Atoi(state["state"].(string))
					if err != nil {
						panic(err)
					}
					moisture = &moistureValue
				case "conductivity":
					conductivityValue, err := strconv.Atoi(state["state"].(string))
					if err != nil {
						// This could be "unavailable"
						panic(err)
					}
					conductivity = &conductivityValue
				case "battery":
					batteryValue, err := strconv.Atoi(state["state"].(string))
					if err != nil {
						panic(err)
					}
					battery = &batteryValue
				case "temperature":
					temperatureValue, err := strconv.ParseFloat(state["state"].(string), 64)
					if err != nil {
						panic(err)
					}
					temperature = &temperatureValue
				case "illuminance":
					illuminanceValue, err := strconv.Atoi(state["state"].(string))
					if err != nil {
						panic(err)
					}
					illuminance = &illuminanceValue
				default:
					panic(fmt.Sprintf("Failure in parsing %s", entityId))
				}
			}
		}
	}
	if illuminance == nil || battery == nil || temperature == nil || moisture == nil || conductivity == nil {
		var msg = fmt.Sprintf("Invalid data in:\n\n%s", string(rawStates))
		panic(msg)
	}
	return SensorStats{
		Illuminance:  *illuminance,
		Battery:      *battery,
		Temperature:  *temperature,
		Moisture:     *moisture,
		Conductivity: *conductivity,
	}
}
