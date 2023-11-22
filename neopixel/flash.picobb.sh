#!/bin/sh
# Canonical build for Raspberry Pi Pico, breadboard prototype.
cd $(dirname "$0")/..
set -ex
exec tinygo flash -target pico -tags picobb -size short -monitor \
     ./neopixel
