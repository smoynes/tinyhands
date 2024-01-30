//go:build microbit_v2

package main

import (
	"machine"
)

var (
	led         = machine.LED_COL_1
	controller  = machine.I2C0
	clockConfig = machine.I2CConfig{
		Frequency: 100e3,
		SDA:       machine.SDA_PIN,
		SCL:       machine.SCL_PIN,
	}
	setClock = false
)
