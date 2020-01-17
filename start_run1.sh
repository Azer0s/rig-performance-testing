#!/bin/bash

set -m

helm install kafka incubator/kafka -f ./kafka/chart/values.yml

while [[ 
    $(kubectl get pods -l app=kafka -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True"
]]; do echo "waiting for pod" && sleep 1; done


helm install rig ./rig/chart

cd src/run1
docker-compose build

chmod +x start_k8s.sh
./start_k8s.sh

while [[ 
    $(kubectl get pods -l app=run1-clients-deployment -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" &&
    $(kubectl get pods -l app=run1-loader-deployment -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" 
]]; do echo "waiting for pod" && sleep 1; done

echo "Starting log collection..."

cd ../..

timeout 60s bash <<EOT
kubectl logs "$(kubectl get pods -l 'app=run1-clients-deployment' -o jsonpath='{.items[0].metadata.name}')" -f > run1.client.log
EOT

helm uninstall rig
helm uninstall kafka