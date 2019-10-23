import sys
sys.path.append("..")
import load as l
import time

def load():
    i = 0
    while True:
        for _ in range(5000):
            l.produce(l.p, "rig", l.payload, "ignore")

        l.print_progress(i, 5000)

        time.sleep(1)

        l.p.flush()
        i += 1