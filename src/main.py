import sys

run = sys.argv[1]

if run == "1":
    print("Run 1")
    from run1 import loader as run1_loader
    run1_loader.load()
elif run == "1_tbd":
    print("Run 1: Load tbd")
    from run1 import loader as run1_loader
    run1_loader.load_tbd()
elif run == "2":
    print("Run 2")
    import loader as run2_loader
    run2_loader.load_10k_include_all()
elif run == "6":
    print("Run 6")
    from run6 import loader as run6_loader
    run6_loader.load_for_multiple_topics(10, 100)
elif run == "7":
    print("Run 7")
    from run7 import loader as run7_loader
    run7_loader.load()
elif run == "7_deliver":
    from run7 import deliver_loader