package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	image_color "image/color"
	"math"
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

const size int32 = 32
const width = 800
const height = 480
const fps = 30
const gradientInterval float64 = 4.0 / fps

var gradientVariable float64

func initRender(config Config) {
	rl.InitWindow(width, height, "Home Assistant")
	if config.Fullscreen {
		rl.ToggleFullscreen()
	}
	rl.SetTargetFPS(fps)
}

func endRender() {
	rl.CloseWindow()
}

func render(stats SensorStats) {
	gradientVariable += gradientInterval
	if gradientVariable > math.Pi*2 {
		gradientVariable = gradientVariable - math.Pi*2
	}
	var gradientSin = math.Sin(gradientVariable)
	green = uint16((gradientSin+1)*(0x7F-0x40)) + 0x40
	// TODO delete to optimize
	if green > 0xFF {
		panic("math error!")
	}
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	var tempText = fmt.Sprintf("Temperature:\t%v\u00B0\nIlluminance:\t%d lx\nBattery:\t%d%%\nConductivity:\t%d\u00B5S/cm\nMoisture:\t%d%%", stats.Temperature, stats.Illuminance, stats.Battery, stats.Conductivity, stats.Moisture)
	rl.DrawText(tempText, 5, 5, size, image_color.RGBA{G: uint8(green), A: 0xFF})
	rl.EndDrawing()
}
