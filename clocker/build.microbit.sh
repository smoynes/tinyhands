#!/bin/sh
# Canonical build for BBC Microbit
cd $(dirname "$0")/..
set -ex

exec tinygo build -target microbit-v2 -programmer cmsis-dap -size short \
     -ldflags="-X 'main.buildTimestamp=\"$(date -Iseconds -u)\"'" \
     -o ./clocker/CLOCKER.uf2 ./clocker
