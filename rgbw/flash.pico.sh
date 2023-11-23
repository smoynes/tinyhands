#!/bin/sh
# Canonical flash for Raspberry Pi Pico "tinyhands"
cd $(dirname "$0")/..
set -ex
exec tinygo flash -target pico -size short \
     -tags picotinyhands \
     -monitor \
     ./rgbw
