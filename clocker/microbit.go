//go:build microbit_v2

package main

import (
	"machine"
)

var (
	led         = machine.LED_COL_1
	controller  = machine.I2C1
	clockConfig = machine.I2CConfig{
		//Frequency: 100e3,
		SDA:       machine.SDA1_PIN,
		SCL:       machine.SCL1_PIN,
	}
	setClock = false
)
