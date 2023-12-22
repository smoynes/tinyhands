//go:build microbit_v2 || pico
// +build microbit_v2 pico

// twinkle is a starlight twinkle experiment on a 0.50m pixel strip.
package main // import "github.com/smoynes/tinyhands/twinkle"

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	ws       ws2812.Device // pixel string device
	leds     []color.RGBA  // pixel array
	state    []uint8
	spectrum []color.RGBA // color palette
)

func init() {
	leds = make([]color.RGBA, numLeds)
	state = make([]uint8, numLeds)

	var colors = [...]color.NRGBA{
		{},                    // Black
		{0xf8, 0xf9, 0xec, 3}, // Starlight
		{0xf2, 0xf9, 0xec, 2},
		{0xf2, 0xf4, 0xec, 1},
	}

	// Adjust brightness
	spectrum = make([]color.RGBA, len(colors))
	for i := 0; i < len(spectrum); i++ {
		r, g, b, a := colors[i].RGBA()
		spectrum[i] = color.RGBA{
			R: uint8(r / 24),
			G: uint8(g / 24),
			B: uint8(b / 24),
			A: uint8(a),
		}
		println(i, spectrum[i].R, spectrum[i].G, spectrum[i].B)
	}
}

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)

	time.Sleep(time.Second)

	println("Running starlight simulator on", machine.Device)
	println("Leds:", len(leds), "Colors:", len(spectrum))

	for i := 0; i < len(state); i++ {
		rnd, err := machine.GetRNG()
		if err != nil {
			panic(err.Error())
		}
		rnd += 1
		rnd %= uint32(len(spectrum))
		state[i] = uint8(rnd)
	}

	led.High()

	const (
		tickInterval = 64 * time.Millisecond
		dt = 64
	)

	for {
		update()
		ws.WriteColors(leds[:])
		time.Sleep(tickInterval)

		eased = ease(dt)
	}
}

func update(easing int) {
	for i := range state {
		curr := state[i]

		if curr == 0 {
			continue
		}

		state[i] = curr
		leds[i] = spectrum[curr]
	}
}
