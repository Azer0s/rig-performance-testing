import sys
sys.path.append("../..")
import load as l

def load_for_multiple_topics(messages_in_k, topics):
    start = l.time.time()
    topic = 1
    for i in range(messages_in_k):
        for _ in range(1000):
            l.produce(l.p, "rig", l.payload, "chatroom_message" + str(topic))

            topic = topic + 1

            if topic > topics:
                topic = 1
        l.p.flush()

        l.print_progress(i)
    end = l.time.time()
    print(end - start)
