FROM ubuntu:18.04
COPY run2/loader.py loader.py
COPY run2/start_loader.sh start_loader.sh
COPY load.py load.py
COPY main.py main.py
RUN apt-get update -y
RUN apt-get install librdkafka-dev python3 python3-pip -y
RUN pip3 install --upgrade pip
RUN pip3 install confluent-kafka
CMD [ "sh", "start_loader.sh" ]