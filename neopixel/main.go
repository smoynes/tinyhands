//go:build microbit_v2 || pico
// +build microbit_v2 pico

// serial is an experiment in microbit serial and terminal I/O.
package main // import "github.com/smoynes/tinyhands/serial"

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var neo machine.Pin = machine.GP0
var ws ws2812.Device
var leds [1]color.RGBA
var spectrum = [...]color.RGBA{
	{148, 0, 211, 255}, // Violet
	{75, 0, 130, 255},  // Indigo
	{0, 0, 255, 255},   // Blue
	{0, 127, 255, 255}, // Blue-Green
	{0, 255, 255, 255}, // Cyan
	{0, 255, 127, 255}, // Aqua
	{0, 255, 0, 255},   // Green
	{127, 255, 0, 255}, // Lime
	{255, 255, 0, 255}, // Yellow
	{255, 127, 0, 255}, // Orange
	{255, 0, 0, 255},   // Red
	{255, 0, 127, 255}, // Pink
	{255, 0, 255, 255}, // Magenta
	{127, 0, 255, 255}, // Purple
}

const spectrumSize = len(spectrum)

const tickInterval = 200 * time.Millisecond

func init() {
	for i := 0; i < len(leds); i += len(spectrum) {
		copy(leds[i:], spectrum[:])
	}
}

func main() {
	println("Running rainbow simulator on", machine.Device)
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)

	last := spectrum[len(spectrum)-1]
	copy(leds[0:len(leds)-1], spectrum[1:len(spectrum)])

	leds[len(leds)-1] = last

	ws.WriteColors(leds[:])
	neo.Set(!neo.Get())

	for {
		time.Sleep(tickInterval)
		first := leds[0]
		copy(leds[:], leds[1:])
		leds[len(leds)-1] = first
		ws.WriteColors(leds[:])
		neo.Set(!neo.Get())
	}
}
