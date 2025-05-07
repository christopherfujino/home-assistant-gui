package main

import (
	"fmt"
	//"math"
	rl "github.com/gen2brain/raylib-go/raylib"
	image_color "image/color"
)

var color = image_color.RGBA{
	G: 0xFF,
	A: 0xFF,
}
var backgroundColor = image_color.RGBA{
	A: 0xFF,
}

// This is only an 8-bit value, but store it in an uint16 so we can detect
// overflow.
var green uint16

const size int32 = 24
const width = 600
const height = 400
const fps = 30
const GRADIENT_INTERVAL = 0x5

func initRender() {
	rl.InitWindow(width, height, "Home Assistant")
	rl.SetTargetFPS(fps)
}

func endRender() {
	rl.CloseWindow()
}

func render(stats SensorStats) {
	//math.Sin()
	green += GRADIENT_INTERVAL
	if green > 0xFF {
		green = 0x20
	}
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	var tempText = fmt.Sprintf("Temperature:\t%v\u00B0\nIlluminance:\t%d lx\nBattery:\t%d%%\nConductivity:\t%d\u00B5S/cm\nMoisture:\t%d%%", stats.Temperature, stats.Illuminance, stats.Battery, stats.Conductivity, stats.Moisture)
	rl.DrawText(tempText, 2, 2, size, image_color.RGBA{G: uint8(green), A: 0xFF})
	rl.EndDrawing()
}
