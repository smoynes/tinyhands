//go:build microbit_v2
// +build microbit_v2

// serial is an experiment in microbit serial and terminal I/O.
package main // import "github.com/smoynes/tinyhands/serial"

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var neo machine.Pin = machine.P1
var ws ws2812.Device
var leds [30]color.RGBA

var colorSpace = color.RGBAModel

var spectrum = [...]color.Color{
	color.NRGBA{148, 0, 211, 255}, // Violet
	color.NRGBA{75, 0, 130, 255},  // Indigo
	color.NRGBA{0, 0, 255, 255},   // Blue
	color.NRGBA{0, 127, 255, 255}, // Blue-Green
	color.NRGBA{0, 255, 255, 255}, // Cyan
	color.NRGBA{0, 255, 127, 255}, // Aqua
	color.NRGBA{0, 255, 0, 255},   // Green
	color.NRGBA{127, 255, 0, 255}, // Lime
	color.NRGBA{255, 255, 0, 255}, // Yellow
	color.NRGBA{255, 127, 0, 255}, // Orange
	color.NRGBA{255, 0, 0, 255},   // Red
	color.NRGBA{255, 0, 127, 255}, // Pink
	color.NRGBA{255, 0, 255, 255}, // Magenta
	color.NRGBA{127, 0, 255, 255}, // Purple
}

func init() {
	for i := 0; i < len(leds); i++ {
		led := spectrum[i%len(spectrum)].(color.RGBA)
		led.A = faders[i]
		c := colorSpace.Convert(led).(color.RGBA)
		leds[i] = c
	}

	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)
}

var faders [255]uint8 = [255]uint8{
	128, 131, 134, 137, 140, 143, 146, 149, 152, 156, 159, 162, 165, 168,
	171, 174, 176, 179, 182, 185, 188, 191, 193, 196, 199, 201, 204, 206,
	209, 211, 213, 216, 218, 220, 222, 224, 226, 228, 230, 232, 234, 235,
	237, 239, 240, 242, 243, 244, 246, 247, 248, 249, 250, 251, 251, 252,
	253, 253, 254, 254, 254, 255, 255, 255, 255, 255, 255, 255, 254, 254,
	253, 253, 252, 252, 251, 250, 249, 248, 247, 246, 245, 244, 242, 241,
	239, 238, 236, 235, 233, 231, 229, 227, 225, 223, 221, 219, 217, 215,
	212, 210, 207, 205, 202, 200, 197, 195, 192, 189, 186, 184, 181, 178,
	175, 172, 169, 166, 163, 160, 157, 154, 151, 148, 145, 142, 138, 135,
	132, 129, 126, 123, 120, 117, 113, 110, 107, 104, 101, 98, 95, 92, 89,
	86, 83, 80, 77, 74, 71, 69, 66, 63, 60, 58, 55, 53, 50, 48, 45, 43, 40,
	38, 36, 34, 32, 30, 28, 26, 24, 22, 20, 19, 17, 16, 14, 13, 11, 10, 9,
	8, 7, 6, 5, 4, 3, 3, 2, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 3,
	4, 4, 5, 6, 7, 8, 9, 11, 12, 13, 15, 16, 18, 20, 21, 23, 25, 27, 29, 31,
	33, 35, 37, 39, 42, 44, 46, 49, 51, 54, 56, 59, 62, 64, 67, 70, 73, 76,
	79, 81, 84, 87, 90, 93, 96, 99, 103, 106, 109, 112, 115, 118, 121, 124,
}

const tickInterval = 200 * time.Millisecond

func main() {
	println("Running rainbow simulator on", machine.Device)

	ws.WriteColors(leds[:])

	flip := neo.Get()
	neo ^= neo
	neo.Set(flip)

	var fadeOffset = 0

	for {
		time.Sleep(tickInterval)

		for i, led := range leds {
			led.A = faders[(i+fadeOffset)%len(faders)]
			c := colorSpace.Convert(led)
			println("led", i, led.R, led.G, led.B, led.A)
			println("color", i, led.R, led.G, led.B)

			//leds[i] = c.(color.RGBA)
		}

		fadeOffset++

		ws.WriteColors(leds[:])
		neo ^= neo
		neo.Set(flip)
	}
}
