# Kedabus

Artifacts from investigating Keda with Service Bus.

## Reproduce the tests

1. Create RG/Service Bus Namespace
2. Create Queue in Namespace called "tasks"
3. Copy Conn String and paste its base64 encoded value in deploy/consumer.yaml after "ConnString"
4. Find a Kube cluster
5. Deploy Keda

    Clone Keda, run:

    ```
    kubectl apply -f deploy/crds/keda.k8s.io_scaledobjects_crd.yaml
    kubectl apply -f deploy/crds/keda.k8s.io_triggerauthentications_crd.yaml
    kubectl apply -f deploy/
    ```

6. Deploy test consumer `kubectl apply -f deploy/`
7. produce messages in the queue

    ```
    go run main.go produce -n 100
    ```