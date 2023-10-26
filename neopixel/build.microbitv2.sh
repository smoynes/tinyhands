#!/bin/sh
# Canonical build for BBC Microbit v2. Run from project root.
set -ex
exec tinygo build -target microbit-v2 -programmer cmsis-dap -size full \
     -o NEOPIXEL.bin ./neopixel
