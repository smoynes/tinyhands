// Clocker is an experiment with an I2C clock module using a DS1307 RTC.

//go:build pico && picotinyhands

package main

import (
	"context"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds1307"
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

var BuildTimestamp string // "2023-11-22T11:31:59-05:00"

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	controller.Configure(clockConfig)
	clock = ds1307.New(controller)
}

func main() {
	time.Sleep(1 * time.Second)
	println("Clocker on", machine.Device)
	println(clock.Address, clock.AddressSRAM, clock.IsOscillatorRunning())

	_, _, done := WithBlinker(context.Background())
	defer done()

	if setClock {
		setClockFromBuildTimestamp()
	}

	for {
		reportTime()
		time.Sleep(1 * time.Second)
	}
}

func setClockFromBuildTimestamp() {
	now, err := time.Parse(time.RFC3339, BuildTimestamp)
	if err != nil {
		panic(err.Error())
	}

	err = clock.SetTime(now)
}

func reportTime() {
	t, err := clock.ReadTime()
	if err != nil {
		panic(err.Error())
	}
	y, mm, d := t.Date()
	h, m, s := t.Clock()

	println(y, mm, d, h, m, s)
}

// WithBlinker starts a task that toggles the LED on two intervals:
//
//   - a short interval, if blinker.Flip has been called during the interval to
//     indicate readiness;
//   - a long interval, in any case, to indicate liveness.
func WithBlinker(ctx context.Context) (context.Context, *blinker, context.CancelFunc) {
	short := time.NewTicker(25 * time.Millisecond)
	long := time.NewTicker(2 * time.Second)

	b := blinker(false)
	go b.blink(ctx, short, long)

	return ctx, &b, func() { short.Stop(); long.Stop() }
}

// Blinker is a boolean that is true if LED blinked in the current short interval.
type blinker bool

func (b *blinker) Flip() { *b = true }

func (b *blinker) blink(ctx context.Context, short *time.Ticker, long *time.Ticker) {
	for {
		for i := uint8(0); i < 10; {
			select {
			case <-ctx.Done():
				return
			case <-long.C:
				led.Set(!led.Get())
				*b = false
			case <-short.C:
				if *b {
					led.Set(!led.Get())
					*b = false
				}
			}

		}

		if *b {
			led.Set(!led.Get())
			*b = false
		}
	}

}
