#!/bin/sh
FAIL=0

set -m

python3 -c "import loader; loader.load_tbd()"

python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
python3 -c "import loader; loader.load()" &
wait

python3 -c "import loader; loader.load_tbd()"