#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
exec tinygo build -target pico -size full \
     -tags picotinyhands \
     -o ./neopixel/twinkle.uf2 ./twinkle
