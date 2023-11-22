#!/bin/sh
# Canonical flash for Raspberry Pi Pico.
cd $(dirname "$0")
set -ex
DATE=$(date -Iseconds -u)
LDFLAGS="-X main.BuildTimestamp=$DATE"
exec tinygo flash -target pico -tags picotinyhands -size short -monitor \
     -ldflags="$LDFLAGS" \
     .
