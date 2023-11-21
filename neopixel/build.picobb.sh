#!/bin/sh
# Canonical build for Raspberry Pi Pico, breadboard prototype.
cd $(dirname "$0")/..
set -ex
exec tinygo build -target pico -tags picobb -size full \
     -o ./neopixel/NEOPIXEL.uf2 ./neopixel
