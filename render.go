package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	image_color "image/color"
)

var color = image_color.RGBA{
	R: 0xFF,
	A: 0xFF,
}

const size int32 = 24
const width = 600
const height = 400

func initRender() {
	rl.InitWindow(width, height, "Home Assistant")
	rl.SetTargetFPS(10)
}

func endRender() {
	rl.CloseWindow()
}

func render(stats SensorStats) {
	rl.BeginDrawing()
	var tempText = fmt.Sprintf("Temperature:\t%v\u00B0\nIlluminance:\t%d lx\nBattery:\t%d%%\nConductivity:\t%d\u00B5S/cm\nMoisture:\t%d%%", stats.Temperature, stats.Illuminance, stats.Battery, stats.Conductivity, stats.Moisture)
	rl.DrawText(tempText, 2, 2, size, color)
	rl.EndDrawing()
}
