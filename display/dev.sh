#!/bin/bash

set -ex

python3 -m venv --system-site-packages ./venv
./venv/bin/pip install displayhatmini
