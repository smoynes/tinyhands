// serial is an experiment in microbit serial and terminal I/O.
package main

import (
	"time"
)

func main() {
	for {
		println("")
		println("Hello!")
		println("... from a Microbit")
		println("... over a USB, serial connection")
		println("... on a Raspberry PI (Model B)")
		println("... in a screen window")
		println("... over an SSH connection")
		println("... down a ethernet cable")
		println("... through a powerline network")
		println("... over a WiFi network")
		println("... to a MacBook Pro! (2019)")
		time.Sleep(2 * time.Second)
	}
}
