#!/bin/sh
# Canonical build for Raspberry Pi Pico.
cd $(dirname "$0")/..
set -ex
exec tinygo flash -target pico -size short \
     -tags picotinyhands -monitor \
     ./twinkle
