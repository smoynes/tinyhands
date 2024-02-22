#!/bin/sh
# Canonical build for BBC Microbit
cd $(dirname "$0")/..
set -ex

exec tinygo flash -target microbit-v2 -programmer cmsis-dap -size short \
     -ldflags="-X 'main.BuildTimestamp=\"$(date -Iseconds -u)\"'" \
     ./clocker
