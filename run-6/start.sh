#!/bin/sh

set -m

cd src
for i in $(seq 1 10); do python3 -c "import loader; loader.load_for_multiple_topics(10, 100)" & bg || true; done
clear