#!/bin/sh
# Canonical build for Raspberry Pi Pico "tinyhands"
cd $(dirname "$0")/..
set -ex
exec tinygo build -target pico -size short \
     -tags picotinyhands \
     -o ./rgbw/RGBW.uf2 ./rgbw
