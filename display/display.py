#!/usr/bin/env python3

import io
import time

import requests
from PIL import Image
from displayhatmini import DisplayHATMini

WIDTH = DisplayHATMini.WIDTH
HEIGHT = DisplayHATMini.HEIGHT

display = DisplayHATMini('', backlight_pwm=True)

# curl http://192.168.2.53/img/snapshot.cgi?size=3 -o image2.jpg

while True:
    url = "http://192.168.2.53/img/snapshot.cgi?size=3"
    req = requests.get(url)
    content = req.content
    image = Image.open(io.BytesIO(content))
    image = image.resize((WIDTH, HEIGHT))
    image = image.convert("RGB")
    print(image)
    display.buffer = image
    display.display()
    time.sleep(5.0)
