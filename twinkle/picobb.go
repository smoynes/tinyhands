//go:build pico && picobb

package main

import "machine"

var (
	neo     = machine.GP2
	led     = machine.LED
	numLeds = 60
)

const rgbw = false
