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

	var (
		start                          = color.NRGBA{10, 10, 10, 255}
		startr, startg, startb, starta = start.RGBA()
		end                            = color.NRGBA{0, 0, 0, 255}
		endr, endg, endb, enda         = end.RGBA()
	)

	const easeSteps = 10
	spectrum = make([]color.RGBA, easeSteps)

	// Ease in-out brightness.
	for i := uint32(0); i < easeSteps; i++ {
		r := int8(startr - i*(startr-endr)/easeSteps)
		g := int8(startg - i*(startg-endg)/easeSteps)
		b := int8(startb - i*(startb-endb)/easeSteps)
		a := int8(starta - i*(starta-enda)/easeSteps)

		spectrum[i] = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
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
		rnd %= uint32(len(spectrum))
		state[i] = uint8(rnd)
	}

	led.High()

	for i := 0; i < len(spectrum); i++ {
		c := spectrum[i]
		println(i, c.R, c.G, c.B, c.A)
	}

	const (
		tickInterval = 100 * time.Millisecond
	)

	for {
		update()
		ws.WriteColors(leds[:])
		time.Sleep(tickInterval)
	}
}

func update() {
	for i := range state {
		curr := state[i]

		if curr == 0 {
			continue
		} else if rnd, err := machine.GetRNG(); err != nil {
			break
		} else {
			incr := (rnd % 3) - 1
			n := curr + uint8(incr)
			if n >= uint8(len(spectrum)) {
				n = curr
			} else if n <= 1 {
				n = 1
			}

			curr = n
		}

		state[i] = curr
		leds[i] = spectrum[curr]
	}
}
