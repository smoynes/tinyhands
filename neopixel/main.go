//go:build microbit_v2 || pico
// +build microbit_v2 pico

// neopixel is an blinkenlights on a strip.
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
	var colorSpectrum = [...]color.NRGBA{
		{148, 0, 211, 3},  // Violet
		{0, 0, 255, 3},    // Blue
		{0, 127, 255, 3},  // Blue-Green
		{0, 255, 255, 3},  // Cyan
		{0, 255, 127, 3},  // Aqua
		{0, 255, 0, 3},    // Green
		{127, 255, 0, 3},  // Lime
		{255, 255, 0, 3},  // Yellow
		{255, 127, 0, 25}, // Orange
		{255, 0, 0, 3},    // Red
		{255, 0, 127, 3},  // Pink
		{255, 0, 255, 3},  // Magenta
		{127, 0, 255, 3},  // Purple

	}

	leds = make([]color.RGBA, numLeds)
	spectrum = make([]color.RGBA, len(colorSpectrum))
	for i := 0; i < len(spectrum); i++ {
		r, g, b, a := colorSpectrum[i].RGBA()
		spectrum[i] = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}
	copy(leds[:], spectrum[:])
}

const tickInterval = 50 * time.Millisecond

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)

	time.Sleep(time.Second)

	println("Running rainbow simulator on", machine.Device)
	led.High()
	println("Leds:", len(leds), "Colors:", len(spectrum))
	for i := 0; i < len(spectrum); i++ {
		c := spectrum[i]
		println("", c.R, c.G, c.B, c.A)
	}

	neo.Set(!neo.Get())

	for {
		time.Sleep(tickInterval)
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
	}
}
