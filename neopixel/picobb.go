//go:build pico && picobb

package main

import "machine"

var (
	neo     = machine.GP16
	led     = machine.LED
	numLeds = 144
)
