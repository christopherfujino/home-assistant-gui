package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func initRender() {
	rl.InitWindow(200, 200, "Home Assistant")
	rl.SetTargetFPS(10)
}

func render(stats SensorStats) {
	fmt.Printf("Temperature:\t%v\u00B0\n", stats.Temperature)
	fmt.Printf("Illuminance:\t%d lx\n", stats.Illuminance)
	fmt.Printf("Battery:\t%d%%\n", stats.Battery)
	fmt.Printf("Conductivity:\t%d\u00B5S/cm\n", stats.Conductivity)
	fmt.Printf("Moisture:\t%d%%\n", stats.Moisture)
}
