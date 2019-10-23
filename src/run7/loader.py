import sys
sys.path.append("..")
import load as l
import time

def load():
    start = time.time()

    for i in range(2):
        timeout = time.time() + 60

        l.produce(l.p, "rig", l.payload, "deliver")

        for _ in range(5000):
            if time.time() > timeout:
                break
            l.produce(l.p, "rig", l.payload, "ignore")

        l.print_progress(i, 5000)

        l.produce(l.p, "rig", l.payload, "deliver")

        left = -1 * (time.time() - timeout)

        if left > 0:
            time.sleep(int(left)) # Sleep the remaining time

        l.p.flush()


    end = l.time.time()
    print(end - start)

load()