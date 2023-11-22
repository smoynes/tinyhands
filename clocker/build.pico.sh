#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
exec tinygo build -target pico -tags picotinyhands -size short \
     -ldflags="-X 'main.buildTimestamp=\"$(date -Iseconds -u)\"'" \
     -o ./clocker/CLOCKER.uf2 ./clocker
