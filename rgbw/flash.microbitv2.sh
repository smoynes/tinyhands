#!/bin/sh
# Canonical flasher for BBC Microbit v2.
cd $(dirname "$0")/..
set -ex
exec tinygo flash -target microbit-v2 -programmer cmsis-dap \
     -monitor ./rgbw
