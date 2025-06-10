package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	image_color "image/color"
)

// light blue
//var color = image_color.RGBA{
//	R: 0x70,
//	G: 0xA4,
//	B: 0xB2,
//	A: 0xFF,
//}
var color = image_color.RGBA{
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

// #000000 black
// #FFFFFF white
// #68372B red
// #70A4B2 light blue
// #6F3D86 purple
// #588D43 green
// #352879 dark blue
// #B8C76F yellow
// #6F4F25 brown
// #433900 dark brown
// #9A6759 light red
// #444444 dark grey
// #6C6C6C mid grey
// #9AD284 light green
// #6C5EB5 mid blue
// #959595 light grey

const spacing = 0
const size int32 = 32
const width = 800
const height = 480
const fps = 30
const gradientInterval float64 = 4.0 / fps

var font rl.Font

func initRender(config Config) {
	rl.InitWindow(width, height, "Home Assistant")
	if config.Fullscreen {
		rl.HideCursor()
		rl.ToggleFullscreen()
	}
	rl.SetTargetFPS(fps)
	// TODO normalize this path
	font = rl.LoadFont("./ignore/C64_TrueType_v1.2.1-STYLE/fonts/C64_Pro_Mono-STYLE.otf")
}

func endRender() {
	rl.CloseWindow()
}

func render(stats SensorStats) {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	const borderDepth = 30
	rl.DrawRectangleLinesEx(rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}, borderDepth, color)
	var tempText = fmt.Sprintf(" *** Plantadore 25 ***\n\nTemperature:\t%vF\nIlluminance:\t%d lx\nConductivity:\t%d S/cm\nMoisture:\t%d%%", stats.Temperature, stats.Illuminance, stats.Conductivity, stats.Moisture)
	rl.DrawTextEx(font, tempText, rl.Vector2{X: borderDepth + 5, Y: float32(borderDepth + font.BaseSize)}, float32(font.BaseSize), spacing, color)
	//rl.DrawText(tempText, 5, 5, size, image_color.RGBA{G: uint8(green), A: 0xFF})
	rl.EndDrawing()
}
