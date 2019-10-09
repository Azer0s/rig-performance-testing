from confluent_kafka import Producer
from sys import stdin
from random import choice
import locale
import uuid
import time

locale.setlocale(locale.LC_ALL, 'en_US')

# This is exactly 1 kB or 1000 bytes
payload = """
{"specversion": "0.2", "type": "???", "id": "###", "data": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
Dui vivamus arcu felis bibendum ut tristique et. 
Quam nulla porttitor massa id neque aliquam vestibulum morbi. 
Vestibulum sed arcu non odio euismod lacinia at quis. 
Ac auctor augue mauris augue neque. 
Purus gravida quis blandit turpis cursus in hac habitasse platea. 
Vulputate eu scelerisque fes imperdiet proin. 
Varius morbi enim nunc faucibus a pellentesque. 
Nec sagittis aliquam aliquam malesuada bibendum arcu. 
Ornare aenean euismod elementum nisi quis eleifend quam adipiscing. 
Amet massa vitae tortor condimentum lacinia quis vel eros.
Nulla aliquet enim tortor at auctor urna nunc id.proin.
Varius morbi enim nunc faucibus a pellentesque.
Nec sagittis aliquam malesuada bibendum arcu
Diam ut venenatis tellus in metus vulputate", "source": "tutorial"}
"""

p = Producer({"bootstrap.servers": "localhost:9092", "message.max.bytes": 2048})

def produce(p, topic, payload, etype):
    p.produce(topic, payload.replace("???", etype).replace("###", str(uuid.uuid1())).encode("utf-8"))

def print_progress(i):
    num = locale.format("%d", ((i + 1) * 1000), grouping=True)
    print(f"Loaded: {num}")

def load():
    start = time.time()

    produce(p, "rig", payload, "to_be_delivered")
    p.flush()

    for i in range(100):
        for _ in range(1000):
            produce(p, "rig", payload, "ignored")
        p.flush()

        print_progress(i)

    produce(p, "rig", payload, "to_be_delivered")
    p.flush()
    
    end = time.time()
    print(end - start)

load()