//go:build pico && picotinyhands

package main

import "machine"

var (
	neo     = machine.GP2
	led     = machine.LED
	numLeds = 144
)

const rgbw = false
