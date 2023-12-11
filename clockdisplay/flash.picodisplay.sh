#!/bin/sh
# Canonical flash for Raspberry Pi Pico.
cd $(dirname "$0")/..
DATE=$(date -Iseconds -u)
LDFLAGS="-X main.BuildTimestamp=$DATE"
set -ex
exec tinygo flash \
     -target pico \
     -size short \
     -tags picodisplay \
     -ldflags="$LDFLAGS" \
     -monitor \
     ./clockdisplay
