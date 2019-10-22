import sys
sys.path.append("../..")
import load as l

def load_10k_include_all():
    start = l.time.time()
    for i in range(10):
        for _ in range(1001):
            l.produce(l.p, "rig", l.payload, "chatroom_message")
        l.p.flush()

        l.print_progress(i)
    end = l.time.time()
    print(end - start)
