#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
DATE=$(date -Iseconds -u)
LDFLAGS="-X main.BuildTimestamp=$DATE"
set -ex
exec tinygo build \
     -target pico \
     -size short \
     -tags picodisplay \
     -ldflags="$LDFLAGS" \
     -o ./clockdisplay/clockdisplay.uf2 \
     ./clockdisplay
