//go:build pico

package main

import (
	"machine"
)

var (
	led         = machine.LED
	controller  = machine.I2C0
	clockConfig = machine.I2CConfig{
		Frequency: 100e3,
		SDA:       machine.GP16,
		SCL:       machine.GP17,
	}
	clock    ds1307.Device
	setClock = false
)
