#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
tinygo flash \
       -target pico \
       -tags picodisplay \
       -programmer raspberrypi-swd \
       -serial uart \
       ./clockdisplay
