// Clockdisplay is an exercise in using an I2C RTC clock module (DS1307) and a
// small LCD display to build a simple clock.

//go:build pico && picodisplay

package main

import (
	"context"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds1307"
	"tinygo.org/x/drivers/st7789"
)

// Pins for Pimoroni Pico Display Pack.
// https://shop.pimoroni.com/products/pico-display-pack
var (
	display  st7789.Device

	LED_R = machine.GP6 // 9
	LED_G = machine.GP7 // 10
	LED_B = machine.GP8 // 11

	SW_A = machine.GP12 // 16
	SW_B = machine.GP13 // 17
	SW_X = machine.GP14 // 19
	SW_Y = machine.GP15 // 20

	LCD_DC   = machine.GP16 // 21
	LCD_CS   = machine.GP17 // 22
	LCD_SCLK = machine.GP18 // 24
	LCD_MOSI = machine.GP19 // 25

	LCD_BL    = machine.GP20  // 26
	LCD_RESET = machine.NoPin // Tied to 30.
)

// Pins for I2C DS1307 RTC
var (
	clock          ds1307.Device

	CLOCK_DA = machine.GP4
	CLOCK_CLK = machine.GP5

	i2c  = machine.I2C0
	clockConfig = machine.I2CConfig{
		Frequency: 100e3,
		SDA:       CLOCK_DA,
		SCL:       CLOCK_CLK,
	}

	setClock       = false
	BuildTimestamp string // "2023-12-11T00:02:40+00:00"

)

func init() {
	LED_R.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_G.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_B.Configure(machine.PinConfig{Mode: machine.PinOutput})

	i2c.Configure(clockConfig)
	clock = ds1307.New(i2c)

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8_000_000,
		SCK:       LCD_SCLK,
		SDO:       LCD_MOSI,
		SDI:       LCD_MOSI,
		Mode:      0,
	})
	display = st7789.New(
		machine.SPI0,
		LCD_RESET,
		LCD_DC,
		LCD_CS,
		LCD_BL,
	)
	display.Configure(st7789.Config{
		Rotation: st7789.NO_ROTATION,
		RowOffset: 80,
		FrameRate: st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
	})
}

func main() {
	time.Sleep(1 * time.Second)

	println("Clockdisplay on", machine.Device)
	width, height := display.Size()
	println(width, height)

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	display.FillScreen(black)

	display.FillRectangle(0, 0, width/2, height/2, white)
	display.FillRectangle(width/2, 0, width/2, height/2, red)
	display.FillRectangle(0, height/2, width/2, height/2, green)
	display.FillRectangle(width/2, height/2, width/2, height/2, blue)
	display.FillRectangle(width/4, height/4, width/2, height/2, black)

	_, _, done := WithBlinker(context.Background())
	defer done()

	if setClock {
		setClockFromBuildTimestamp()
	}

	for {
		//reportTime()
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
				LED_R.Set(!LED_R.Get())
				*b = false
			case <-short.C:
				if *b {
					LED_R.Set(!LED_R.Get())
					*b = false
				}
			}

		}

		if *b {
			LED_R.Set(!LED_R.Get())
			*b = false
		}
	}

}
