#!/bin/sh
# Canonical flasher for pico
cd $(dirname "$0")/..
set -ex
exec tinygo flash -target pico -size short \
     -monitor ./ringer
