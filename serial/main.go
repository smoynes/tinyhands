//go:build microbit_v2 || pico
// +build microbit_v2 pico

// serial is an experiment in microbit and pico serial and terminal I/O.
package main // import "github.com/smoynes/tinyhands/serial"

import (
	"time"
)

const tickInterval = 2 * time.Second

func main() {
	tickTime := time.Now()

	for {
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

		time.Sleep(tickInterval)
	}
}
