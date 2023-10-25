//go:build microbit_v2
// +build microbit_v2

// serial is an experiment in microbit serial and terminal I/O.
package main // import "github.com/smoynes/tinyhands/serial"

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

var matrix microbitmatrix.Device

func init() {
	matrix = microbitmatrix.New()
	matrix.Configure(microbitmatrix.Config{
		Rotation: microbitmatrix.RotationNormal,
	})
}

const refreshInterval = 16 * time.Microsecond
const tickInterval = 2 * time.Second
const brightLevel = int8(255 / 9)

func main() {
	matrix.ClearDisplay()

	tickTime := time.Now()
	refreshTime := tickTime

	var ledx, ledy, ledIndex int16

	for {
		if time.Since(refreshTime) > refreshInterval {
			matrix.Display()
			refreshTime = time.Now()
			continue
		}

		if time.Since(tickTime) <= tickInterval {
			continue
		}

		tickTime = time.Now()

		println("Hello there! I blink. 😆")
		println("... from a Microbit\n")
		println(" ... over a USB, serial connection\n")
		println("  ... on a Raspberry PI (Model B)\n")
		println("   ... in a screen window\n")
		println("    ... over an SSH connection\n")
		println("     ... down a ethernet cable\n")
		println("      ... through a powerline network\n")
		println("       ... over a WiFi network\n")
		println("        ... to a MacBook Pro! (2019)\n")

		matrix.ClearDisplay()
		for i := uint8(0); i < 9; i++ {
			bright := 0xff - i*uint8(brightLevel)
			colour := color.RGBA{0xff, 0xff, 0xff, bright}
			ledx, ledy =
				(ledIndex+int16(i))/5,
				(ledIndex+int16(i))%5
			matrix.SetPixel(ledx, ledy, colour)

		}
		matrix.Display()
		refreshTime = time.Now()

		ledIndex++
		ledIndex %= 25

	}
}
