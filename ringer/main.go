//go:build pico

// ringserial is a demo using serial input and a neopixel ring as output.
package main

import (
	"context"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	ring machine.Pin = machine.GP15
	led  machine.Pin = machine.LED
)

var (
	serial machine.Serialer = machine.Serial
)

func main() {
	led.High()

	led.Configure(machine.PinConfig{})
	machine.InitSerial()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	time.Sleep(1 * time.Second)
	speed := machine.CPUFrequency() / machine.MHz
	println("Welcome to", machine.Device,
		"running at", speed, "MHz")
	println("WS2812", ring)
	ctx, cancel = WithWatchdog(ctx)
	defer cancel()

	var blink *blinker
	ctx, blink, cancel = WithBlinker(ctx)
	defer cancel()

	ctx, cancel = WithMonitor(ctx, blink)
	defer cancel()

	for {
		if !waitForEvents(ctx) {
			break
		}
	}

	if ctx.Err() != nil {
		serial.Write([]byte("ERROR: "))
		serial.Write([]byte(ctx.Err().Error()))
		serial.WriteByte('\n')
	}
}

func waitForEvents(ctx context.Context) bool {
	if ctx.Err() != nil {
		return false
	}

	time.Sleep(10 * time.Second)

	return ctx.Err() == nil
}

// WithWatchdog starts a task that updates the watchdog timer periodically.
func WithWatchdog(ctx context.Context) (context.Context, context.CancelFunc) {
	var watchdogMillis uint32 = 1_000

	dog := machine.WatchdogConfig{TimeoutMillis: watchdogMillis}
	machine.Watchdog.Configure(dog)
	tick := time.NewTicker(
		time.Duration(watchdogMillis-100) * time.Millisecond,
	)

	go watchdog(ctx, tick)

	return ctx, tick.Stop
}

func watchdog(ctx context.Context, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			machine.Watchdog.Update()
		case <-ctx.Done():
			return
		}
	}

}

// WithBlinker starts a task periodically blink a number of LEDs:
//
//   - on a short interval, advance an LED on a ring.
//
// /  - on a long interval, toggle the board LED to indicate liveness.
func WithBlinker(ctx context.Context) (context.Context, *blinker, context.CancelFunc) {
	short := time.NewTicker(60 * time.Millisecond)
	long := time.NewTicker(2 * time.Second)

	neo := ws2812.New(ring)

	b := blinker{
		leds: make([]color.RGBA, 12),
		neo:  neo,
	}
	n := copy(b.leds, []color.RGBA{
		color.RGBA{255, 0, 0, 0},
		color.RGBA{0, 255, 0, 0},
		color.RGBA{0, 0, 255, 0},
	})
	println(n)

	go b.blink(ctx, short, long)

	return ctx, &b, func() { short.Stop(); long.Stop() }
}

// blinker holds the task state
type blinker struct {
	flipped bool
	index   uint8
	neo     ws2812.Device
	leds    []color.RGBA
}

func (b *blinker) Flip() { b.flipped = true }

func (b *blinker) blink(ctx context.Context, short *time.Ticker, long *time.Ticker) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-long.C:
			led.Set(!led.Get())
		case <-short.C:
			if b.flipped {
				println("colors", len(b.leds))
				b.neo.WriteColors(b.leds)
			}

			b.flipped = false
		}
	}

}

// WithMonitor starts a task that blinks the when a byte is read from the
// default UART and buffered.
func WithMonitor(ctx context.Context, blinker *blinker) (context.Context, context.CancelFunc) {
	m := monitor{serial, blinker}

	go m.monitor(ctx)

	return ctx, func() {}
}

type monitor struct {
	uart  machine.Serialer
	blink *blinker
}

// monitor reads bytes from uart
func (m monitor) monitor(ctx context.Context) {
	for {
		if m.uart.Buffered() > 0 {
			byte, err := m.uart.ReadByte()
			if err != nil {
				panic(err.Error())
			}

			m.blink.Flip()
			println("read", byte)
		}

		time.Sleep(time.Millisecond) // slow
	}
}
