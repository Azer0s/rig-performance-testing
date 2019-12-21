#!/bin/bash

# VERY TEMPORARY. SOMETHING IS KINDA BROKEN.
# WILL PROBABLY NEED TO MOVE HTTP HANDLER TO AFTER I HAVE CREATED THE CONNECTIONS

sleep 45

while true; do
    status=$(curl -s "clients:9999")

    if [[ $status == *"OK"* ]]; then
        break
    else
        echo "."
    fi

    sleep 1
done

echo "Starting loader..."