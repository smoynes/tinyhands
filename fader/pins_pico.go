//go:build pico
// +build pico

package main

import (
	"machine"
)

var (
	neo machine.Pin = machine.GP15
)
