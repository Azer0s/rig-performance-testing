#!/bin/sh
FAIL=0

set -m

sleep 45

python3 -c 'import loader as run1_loader; run1_loader.load_tbd()'

python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
python3 -c 'import loader as run1_loader; run1_loader.load()' &
wait

python3 -c 'import loader as run1_loader; run1_loader.load_tbd()'
