#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
exec tinygo build -target pico -size full \
     -o ./neopixel/NEOPIXEL.uf2 ./neopixel
