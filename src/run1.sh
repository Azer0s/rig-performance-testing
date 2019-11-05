#!/bin/sh
FAIL=0

set -m

python3 main.py 1_tbd

python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
python3 main.py 1 &
wait

python3 main.py 1_tbd
