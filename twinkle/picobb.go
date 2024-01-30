//go:build pico && picobb

package main

import "machine"

var (
	neo     = machine.GP12
	led     = machine.LED
	numLeds = 60
)

const rgbw = false
