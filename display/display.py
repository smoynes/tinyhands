#!/usr/bin/env python3

import time

from PIL import Image, ImageDraw, ImageFont
from displayhatmini import DisplayHATMini

WIDTH = DisplayHATMini.WIDTH
HEIGHT = DisplayHATMini.HEIGHT

font = ImageFont.load_default()
image = Image.new("RGB", (WIDTH, HEIGHT))
draw = ImageDraw.Draw(image)
display = DisplayHATMini(image, backlight_pwm=True)

while True:
    display.display()
    time.sleep(1.0 / 30)
