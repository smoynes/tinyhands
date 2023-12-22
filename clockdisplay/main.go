// Clockdisplay is an exercise in using an I2C RTC clock module (DS1307) and a
// small LCD display to build a simple clock.

//go:build pico && picodisplay

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds1307"
	"tinygo.org/x/drivers/st7789"
)

func main() {
	time.Sleep(2 * time.Second)

	println("Clockdisplay on", machine.Device)

	setupLED()
	//setupClock()
	setupDisplay()

	time.Sleep(500 * time.Millisecond)

	width, height := display.Size()
	println(width, height)

	drawDisplay()
}

func setupLED() {
	LED_R.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_G.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_B.Configure(machine.PinConfig{Mode: machine.PinOutput})

	LED_R.Set(false)
	LED_G.Set(false)
	LED_B.Set(false)
}

// Pins for Pimoroni Pico Display Pack.
// https://shop.pimoroni.com/products/pico-display-pack
var (
	display st7789.Device

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
	LCD_MISO = machine.NoPin

	LCD_BL    = machine.GP20  // 26
	LCD_RESET = machine.NoPin // Tied to 30.
)

func setupDisplay() {
	println("Configuring SPI")
	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8_000_000,
		Mode:      0,
	})
	if err != nil {
		panic(err.Error())
	}
	display = st7789.New(
		machine.SPI0,
		LCD_RESET,
		LCD_DC,
		LCD_CS,
		LCD_BL,
	)
	display.Configure(st7789.Config{
		Height:     135,
		Width:      240,
		RowOffset:  0,
		FrameRate:  st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
	})
	println("done")
}

func drawDisplay() {
	white := color.RGBA{255, 255, 255, 255}
	/*red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}*/
	black := color.RGBA{0, 0, 0, 255}

	display.FillScreen(black)

	width, height := int16(240), int16(135)
	display.FillRectangle(0, 0, width-5, height-5, white)
}

// Pins for I2C DS1307 RTC
var (
	clock ds1307.Device

	CLOCK_DA  = machine.GP12
	CLOCK_CLK = machine.GP13

	i2c         = machine.I2C0
	clockConfig = machine.I2CConfig{
		Frequency: 100e3,
		SDA:       CLOCK_DA,
		SCL:       CLOCK_CLK,
	}

	setClock       = false
	BuildTimestamp string // "2023-12-11T00:02:40+00:00"

)

func setupClock() {
	err := i2c.Configure(clockConfig)
	if err != nil {
		panic(err.Error())
	}
	clock = ds1307.New(i2c)
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
