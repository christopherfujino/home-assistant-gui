package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	image_color "image/color"
)

// mid blue
var foregroundColor = image_color.RGBA{
	R: 0x65,
	G: 0x5E,
	B: 0xB5,
	A: 0xFF,
}

// dark blue
var backgroundColor = image_color.RGBA{
	R: 0x35,
	G: 0x28,
	B: 0x79,
	A: 0xFF,
}

const spacing = 0
const size int32 = 32
const width = 800
const height = 480
const fps = 30
const gradientInterval float64 = 4.0 / fps
const borderDepth = 30

var font rl.Font

func initRender(config Config) {
	rl.InitWindow(width, height, "Home Assistant")
	if config.Fullscreen {
		rl.HideCursor()
		rl.ToggleFullscreen()
	}
	rl.SetTargetFPS(fps)
	font = rl.LoadFont("./ignore/C64_TrueType_v1.2.1-STYLE/fonts/C64_Pro_Mono-STYLE.otf")
}

func endRender() {
	rl.UnloadFont(font)
	rl.CloseWindow()
}

func render(stats Stats) {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	rl.DrawRectangleLinesEx(rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}, borderDepth, foregroundColor)
	var tempText = fmt.Sprintf(" ** Planttadore 25 **\n\nTemperature:\t%vF\nIlluminance:\t%d lx\nConductivity:\t%d S/cm\nMoisture:\t%d%%\nCPU Temp:\t%.1fC", stats.Temperature, stats.Illuminance, stats.Conductivity, stats.Moisture, stats.CPUTemp)
	rl.DrawTextEx(font, tempText, rl.Vector2{X: borderDepth + 5, Y: float32(borderDepth + font.BaseSize)}, float32(font.BaseSize), spacing, foregroundColor)
	rl.EndDrawing()
}
