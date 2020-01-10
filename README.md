# Giant Swarm TODO

This is a simple TODO list demo application that we use to demonstrate some possibilities
enabled by a cloud native stack. This demo is best suited for testing Loki and Linkerd2.

## Requirements

There are no real requirements and you can install the application as-is.

Still, if you want to use it with Loki and Linkerd2, please see instructions below.

## Installation

### Dependencies

To install loki, run:

```bash
helm install --namespace loki -n loki --version 0.2.0 giantswarm-playground/loki-stack-app
```

To install Linkerd2, run:

```bash
helm install --namespace linkerd -n linkerd -f ./helm/configs/linkerd.yaml --version 0.2.1 giantswarm-playground/linkerd2-app
```

### Deployment

This application needs some parameters to setup the MySQL database used as data store. By default it runs
with no persistent storage for mysql. You can start with the default configuration included in the repo:

```text
helm install --name gs-todo --namespace todo --version 0.2.2 giantswarm-playground/giantswarm-todo-app
```

## Configuration

| Variable                  | Default                  | Description                                               |
| ------------------------- | ------------------------ | --------------------------------------------------------- |
| clusterDomain             | cluster.local            | your domain configured for Kubernetes cluster             |
| todomanagerReplicaCount   | 3                        | How many replicas to create for todomanager               |
| apiserverReplicaCount     | 3                        | How many replicas to create for apiserver                 |
| resources                 | cpu: 100m, memory: 128Mi | Kubernetes limits used for todomanager and apiserver pods |
| linkerdEnabled            | true                     | If integration with linkerd should be enabled             |
| linkerdNamespace          | linkerd                  | Namespace where linkerd is deployed (only if enabled)     |
| apiserverServiceType      | ClusterIP                | Service type for apiserver                                |
| todomanagerServiceType    | ClusterIP                | Service type for todomanager                              |
| ingress.enabled           | false                    | Should Ingress be configured for the application          |
| mysql.persistence.enabled | false                    | Should MySQL use persistent storage for data              |

## Compatibility

Tested on Giant Swarm release 10.1.0 on AWS with Kubernetes 1.15.5.
