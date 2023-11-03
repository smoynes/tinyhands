#!/bin/sh
# Canonical build for BBC Microbit v2.
cd $(dirname $0)/..
set -ex
exec tinygo build \
     -target microbit-v2-s113v7 \
     -programmer cmsis-dap \
     -o BLUETOOTH.bin \
     ./bluetooth
