// uart is an exercise in using the Universal Asynchronous Receive Transmit
// interface and also concurrency.
package main

import (
	"context"
	"machine"
	"time"
)

var (
	led                      = machine.LED
	serial  machine.Serialer = machine.Serial // The primary UART.
	secrets *machine.UART    = machine.UART1  // The secondary UART.
)

var blink bool

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

	println("Configured UARTs:")
	println(" * UART0\n * UART1")

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
			println("woof")
			machine.Watchdog.Update()
		case <-ctx.Done():
			return
		}
	}

}

// WithBlinker starts a task that toggles the LED on two intervals:
//
//   - a short interval, if blinker.Flip has been called during the interval to
//     indicate readiness;
//   - a long interval, in any case, to indicate liveness.
func WithBlinker(ctx context.Context) (context.Context, *blinker, context.CancelFunc) {
	short := time.NewTicker(100 * time.Millisecond)
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

// WithMonitor starts a task that blinks the when a byte is read from the
// default UART and buffered.
func WithMonitor(ctx context.Context, blinker *blinker) (context.Context, context.CancelFunc) {
	m := monitor{serial, blinker}
	input := make(chan rune, 2)

	go m.monitor(ctx, input)

	return ctx, func() {
		serial.Write([]byte("\n\nBye\n\n"))
		close(input)
	}
}

type monitor struct {
	uart  machine.Serialer
	blink *blinker
}

// monitor reads bytes from uart
func (m monitor) monitor(ctx context.Context, read chan<- rune) {
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
