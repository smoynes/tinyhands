#!/bin/sh
# Canonical build for BBC Microbit v2.
exec tinygo build -target microbit-v2 -o serial.elf
