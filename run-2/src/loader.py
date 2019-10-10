from confluent_kafka import Producer
from confluent_kafka.admin import AdminClient, NewTopic
from confluent_kafka import KafkaError
from sys import stdin
from random import choice
from string import ascii_lowercase
import locale
import uuid
import time
import asyncio

locale.setlocale(locale.LC_ALL, 'en_US')

a = AdminClient({"bootstrap.servers": "localhost:9092"})

def delete_topic():
    for _, f in a.delete_topics(["rig"]).items():
        while not f.done():
            print("", end="")
        print("Topic deleted...")

def recreate_topic():
    for _, f in a.create_topics([NewTopic("rig", 8, replication_factor=1)]).items():
        while not f.done():
            print("", end="")
        print("Topic recreated...")


def clear_topic():
    delete_topic()
    recreate_topic()

# This is exactly 1 kB or 1000 bytes
payload = """
{"specversion": "0.2", "type": "chatroom_message", "id": "###", "data": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
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

def produce(p, topic, payload, header):
    p.produce(topic, payload.replace("###", str(uuid.uuid1())).encode("utf-8"), headers=header)

def print_progress(i):
    num = locale.format("%d", ((i + 1) * 1000), grouping=True)
    print(f"Loaded: {num}")

def load_10k_include_all(header={"eventType": "to_be_delivered"}):
    start = time.time()
    for i in range(10):
        for _ in range(1001):
            produce(p, "rig", payload, header)
        p.flush()

        print_progress(i)
    end = time.time()
    print(end - start)
