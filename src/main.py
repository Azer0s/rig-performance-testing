from runs.run1 import loader as run1_loader
from runs.run2 import loader as run2_loader
from runs.run6 import loader as run6_loader
from runs.run7 import loader as run7_loader

import sys

run = sys.argv[1]

if run == "1":
    print("Run 1")
    run1_loader.load()
elif run == "1_tbd":
    print("Run 1: Load tbd")
    run1_loader.load_tbd()
elif run == "2":
    print("Run 2")
    run2_loader.load_10k_include_all()
elif run == "6":
    print("Run 6")
    run6_loader.load_for_multiple_topics(10, 100)
elif run == "7":
    print("Run 7")
    run7_loader.load()
elif run == "7_deliver":
    from runs.run7 import deliver_loader