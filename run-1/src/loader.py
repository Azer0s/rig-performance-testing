import sys
sys.path.append("../..")
import load as l

def load():
    start = l.time.time()

    l.produce(l.p, "rig", l.payload, "to_be_delivered")
    l.p.flush()

    for i in range(1000):
        for _ in range(1000):
            l.produce(l.p, "rig", l.payload, "ignored")
        l.p.flush()

        l.print_progress(i)

    l.produce(l.p, "rig", l.payload, "to_be_delivered")
    l.p.flush()
    
    end = l.time.time()
    print(end - start)

load()