//go:build microbit_v2 || pico
// +build microbit_v2 pico

// twinkle is a starlight twinkle experiment on a 0.50m pixel strip.
package main // import "github.com/smoynes/tinyhands/serial"

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var ws ws2812.Device
var leds []color.RGBA
var spectrum []color.RGBA

func init() {
	var colors = [...]color.NRGBA{
		{0xf8, 0xf9, 0xec, 3}, // Starlight
		{0xf2, 0xf9, 0xec, 2},
		{0xF9, 0xF4, 0xEC, 1},
	}

	leds = make([]color.RGBA, numLeds)
	spectrum = make([]color.RGBA, len(colors))
	for i := 0; i < len(spectrum); i++ {
		r, g, b, a := colors[i].RGBA()
		spectrum[i] = color.RGBA{
			R: uint8(r / 20),
			G: uint8(g / 20),
			B: uint8(b / 20),
			A: uint8(a),
		}
	}
	for i := 0; i < numLeds; i += numLeds / 4 {
		copy(leds[i:], spectrum[:])
	}
}

const tickInterval = 22 * time.Millisecond

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)

	time.Sleep(time.Second)

	println("Running rainbow simulator on", machine.Device)
	println("Leds:", len(leds), "Colors:", len(spectrum))
	for i := 0; i < len(leds); i++ {
		c := leds[i]
		if !(c.R == 0 && c.G == 0 && c.B == 0) {
			println(i, c.R, c.G, c.B, c.A)
		}
	}

	led.High()

	neo.Set(!neo.Get())

	for {
		first := leds[0]
		copy(leds[:], leds[1:])
		leds[len(leds)-1] = first

		if rgbw {
			for i := 0; i < len(leds); i++ {
				c := leds[i]
				ws.WriteByte(c.R)
				ws.WriteByte(c.B)
				ws.WriteByte(c.G)
				ws.WriteByte(1)
			}
		} else {
			ws.WriteColors(leds[:])

		}

		time.Sleep(tickInterval)
	}
}
