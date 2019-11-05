import sys
from .. import load as l
import time

while True:
    l.produce(l.p, "rig", l.payload, "deliver")
    time.sleep(1)
    l.p.flush()