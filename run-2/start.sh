#!/bin/sh

cd src
#python3 -c "import loader; loader.clear_topic()"
for i in $(seq 1 10); do python3 -c "import loader; loader.load_10k_include_all()" & bg || true; done
clear