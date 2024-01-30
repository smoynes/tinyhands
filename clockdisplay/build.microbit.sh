#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex

tinygo build \
     -target microbitv2 \
     -size short \
     -tags microbit \
     ./clockdisplay
