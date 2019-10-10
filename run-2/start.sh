#!/bin/sh

set -m

cd src
for i in $(seq 1 10); do python3 -c "import loader; loader.load_10k_include_all()" & bg || true; done
clear