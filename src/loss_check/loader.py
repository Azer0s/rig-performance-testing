import sys
import load as l
import time

payload = """
{"specversion": "0.2", "type": "???", "id": "###", "data": "+++", "source": "tutorial"}
"""

payload_done = """
{"specversion": "0.2", "type": "???", "id": "###", "data": "DONE", "source": "tutorial"}
"""

def load():
    i = 0
    for j in range(10):
        for k in range(1000):
            l.produce(l.p, "rig", payload.replace("+++", str((j * 1000) + k)), "deliver")

        l.print_progress(i)

        time.sleep(1)

        l.p.flush()
        i += 1

    for _ in range(1000):
        l.produce(l.p, "rig", payload_done, "deliver")

    l.p.flush()

load()