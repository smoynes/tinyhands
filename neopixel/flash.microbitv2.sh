#!/bin/sh
# Canonical flash for BBC Microbit v2. Run from project root.
set -ex
exec tinygo flash -target microbit-v2 -programmer cmsis-dap \
     -size full -monitor \
     ./neopixel
