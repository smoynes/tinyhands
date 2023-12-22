#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
# For bootsel
tinygo build \
     -target pico \
     -size short \
     -tags picodisplay \
     -o ./clockdisplay/clockdisplay.uf2 \
     ./clockdisplay

# For openocd
tinygo build \
       -target pico \
       -tags picodisplay \
       -o ./clockdisplay/clockdisplay.elf \
       ./clockdisplay
