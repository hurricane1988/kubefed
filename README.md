[![Github Actions](https://github.com/kubernetes-sigs/kubefed/actions/workflows/test-and-push.yml/badge.svg?branch=master)](https://github.com/kubernetes-sigs/kubefed/actions?query=branch%3Amaster "Github Actions")
[![Go Report Card](https://goreportcard.com/badge/github.com/kubernetes-sigs/kubefed)](https://goreportcard.com/report/github.com/kubernetes-sigs/kubefed)
[![Image Repository on Quay](https://quay.io/repository/kubernetes-multicluster/kubefed/status "Image Repository on Quay")](https://quay.io/repository/kubernetes-multicluster/kubefed)
[![LICENSE](https://img.shields.io/badge/license-apache2.0-green.svg)](https://github.com/kubernetes-sigs/kubefed/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/release/kubernetes-sigs/kubefed/all.svg)](https://github.com/kubernetes-sigs/kubefed/releases "KubeFed latest release")

---
## Running status
```shell
╭━╮╭━╮╱╱╭╮╭╮╱╱╱╱╭━━━┳╮╱╱╱╱╱╱╭╮╱╱╱╱╱╱╱╭╮╭━╮╱╱╭╮╱╱╱╱╱╭━╮╱╱╱╱╭╮
┃┃╰╯┃┃╱╱┃┣╯╰╮╱╱╱┃╭━╮┃┃╱╱╱╱╱╭╯╰╮╱╱╱╱╱╱┃┃┃╭╯╱╱┃┃╱╱╱╱╱┃╭╯╱╱╱╱┃┃
┃╭╮╭╮┣╮╭┫┣╮╭╋╮╱╱┃┃╱╰┫┃╭╮╭┳━┻╮╭╋━━┳━╮╱┃╰╯╯╭╮╭┫╰━┳━━┳╯╰┳━━┳━╯┃
┃┃┃┃┃┃┃┃┃┃┃┃┣╋━━┫┃╱╭┫┃┃┃┃┃━━┫┃┃┃━┫╭┻━┫╭╮┃┃┃┃┃╭╮┃┃━╋╮╭┫┃━┫╭╮┃
┃┃┃┃┃┃╰╯┃╰┫╰┫┣━━┫╰━╯┃╰┫╰╯┣━━┃╰┫┃━┫┣━━┫┃┃╰┫╰╯┃╰╯┃┃━┫┃┃┃┃━┫╰╯┃
╰╯╰╯╰┻━━┻━┻━┻╯╱╱╰━━━┻━┻━━┻━━┻━┻━━┻╯╱╱╰╯╰━┻━━┻━━┻━━╯╰╯╰━━┻━━╯

+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+--------------+---------------+--------------+
| COMMUNITY  | AUTHOR       | VERSION | GIT COMMIT                               | BUILD DATE           | GO VERSION | COMPILER | PLATFORM     | RUNTIME CORES | TOTAL MEMORY |
+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+--------------+---------------+--------------+
| CodeFuture | Jianping Niu | v1.1.0  | 137cff22e5580fda7851e3b86300280c121d2cbf | 2026-03-25T02:38:58Z | go1.26.1   | gc       | darwin/arm64 | 10 cores      | 9158 KB      |
+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+--------------+---------------+--------------+
The KubeFed controller manager runs a bunch of controllers
which watch KubeFed CRDs and the corresponding resources in
member clusters and do the necessary reconciliation
````
---
## Project layout
```shell
.
├── CHANGELOG
│   ├── CHANGELOG-v1.0.1.md
│   ├── CHANGELOG-v1.0.2.md
│   └── CHANGELOG-v1.1.0.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── OWNERS
├── PROJECT
├── README.md
├── SECURITY.md
├── SECURITY_CONTACTS
│   ├── controller-manager
│   ├── hyperfed
│   ├── kubefedctl
│   └── webhook
├── build
│   └── kubefed
├── charts
│   ├── index.yaml
│   ├── kubefed
│   └── service-monitor.yaml
├── cmd
│   ├── controller-manager
│   ├── hyperfed
│   ├── kubefedctl
│   └── webhook
├── code-of-conduct.md
├── config
│   ├── crds
│   ├── enabletypedirectives
│   └── kubefedconfig.yaml
├── docs
│   ├── OWNERS
│   ├── cluster-registration.md
│   ├── concepts.md
│   ├── development.md
│   ├── environments
│   ├── images
│   ├── installation.md
│   ├── keps
│   ├── project-layout.md
│   ├── releasing.md
│   └── userguide.md
├── example
│   ├── config
│   ├── sample
│   └── scheduling
├── go.mod
├── go.sum
├── hack
│   ├── boilerplate.go.txt
│   ├── doc.go
│   ├── verify-docfiles.sh
│   ├── verify-errpkg.sh
│   └── verify-klog.sh
├── pkg
│   ├── apis
│   ├── client
│   ├── constants
│   ├── controller
│   ├── doc.go
│   ├── features
│   ├── kubefedctl
│   ├── metrics
│   ├── schedulingtypes
│   └── version
├── scripts
│   ├── build-release-artifacts.sh
│   ├── build-release.sh
│   ├── check-directive-fixtures.sh
│   ├── create-clusters.sh
│   ├── create-gh-release.sh
│   ├── delete-clusters.sh
│   ├── delete-kubefed.sh
│   ├── deploy-federated-nginx.sh
│   ├── deploy-kubefed-latest.sh
│   ├── deploy-kubefed.sh
│   ├── download-binaries.sh
│   ├── download-e2e-binaries.sh
│   ├── fix-ca-for-k3s.sh
│   ├── fix-joined-kind-clusters.sh
│   ├── pre-commit.sh
│   ├── sync-up-helm-chart.sh
│   ├── update-bindata.sh
│   └── util.sh
├── staticcheck.conf
├── test
│   ├── common
│   └── e2e
├── third-party
│   └── k8s.io
└── tools
    ├── go.mod
    ├── go.sum
    └── tools.go
```
---
## Project Usage
```shell
Usage:
  controller-manager [flags]

Flags:
      --add_dir_header                         If true, adds the file directory to the header of the log messages
      --alsologtostderr                        log to standard error as well as files (no effect when -logtostderr=true)
      --alsologtostderrthreshold severity      logs at or above this threshold go to stderr when -alsologtostderr=true (no effect when -logtostderr=true)
      --cluster-health-check-period duration   How often to check the health of cluster. (default 10s)
      --healthz-addr string                    The address the healthz endpoint binds to. (default ":8080")
  -h, --help                                   help for controller-manager
      --kubeconfig string                      Path to a kubeconfig. Only required if out-of-cluster.
      --kubefed-config string                  Path to a KubeFedConfig yaml file. Test only.
      --kubefed-namespace string               The namespace the KubeFed control plane is deployed in.
      --leader-elect-lease-duration duration   The maximum duration that a leader can be stopped before it is replaced by another candidate. (default 15s)
      --leader-elect-renew-deadline duration   The interval between attempts by the acting master to renew a leadership slot before it stops leading. (default 10s)
      --leader-elect-resource-lock string      The type of resource object that will be used to lock during leader election. (default "leases")
      --leader-elect-retry-period duration     The duration the clients should wait between attempting acquisition and renewal of a leadership. (default 5s)
      --legacy_stderr_threshold_behavior       If true, stderrthreshold is ignored when logtostderr=true (legacy behavior). If false, stderrthreshold is honored even when logtostderr=true (default true)
      --log_backtrace_at traceLocation         when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                         If non-empty, write log files in this directory (no effect when -logtostderr=true)
      --log_file string                        If non-empty, use this log file (no effect when -logtostderr=true)
      --log_file_max_size uint                 Defines the maximum size a log file can grow to (no effect when -logtostderr=true). Unit is megabytes. If the value is 0, the maximum file size is unlimited. (default 1800)
      --logtostderr                            log to standard error instead of files (default true)
      --master string                          The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
      --metrics-addr string                    The address the metric endpoint binds to. (default ":9090")
      --one_output                             If true, only write logs to their native severity level (vs also writing to each lower severity level; no effect when -logtostderr=true)
      --rest-config-burst int                  Maximum burst for throttle to the api-server from this client. (default 200)
      --rest-config-qps float32                Maximum QPS to the api-server from this client. (default 100)
      --skip_headers                           If true, avoid header prefixes in the log messages
      --skip_log_headers                       If true, avoid headers when opening log files (no effect when -logtostderr=true)
      --stderrthreshold severity               logs at or above this threshold go to stderr when writing to files and stderr (no effect when -logtostderr=true or -alsologtostderr=true unless -legacy_stderr_threshold_behavior=false) (default 2)
  -v, --v Level                                number for the log level verbosity
      --version                                Prints the Version info of controller-manager.
      --vmodule moduleSpec                     comma-separated list of pattern=N settings for file-filtered logging
```
# Kubernetes Cluster Federation

Kubernetes Cluster Federation (KubeFed for short) allows you to coordinate the
configuration of multiple Kubernetes clusters from a single set of APIs in a
hosting cluster. KubeFed aims to provide mechanisms for expressing which
clusters should have their configuration managed and what that configuration
should be. The mechanisms that KubeFed provides are intentionally low-level, and
intended to be foundational for more complex multicluster use cases such as
deploying multi-geo applications and disaster recovery.

KubeFed is currently **beta**.

## Concepts

<p align="center"><img src="docs/images/concepts.png" width="711"></p>

KubeFed is configured with two types of information:

- **Type configuration** declares which API types KubeFed should handle
- **Cluster configuration** declares which clusters KubeFed should target

**Propagation** refers to the mechanism that distributes resources to federated
clusters.

Type configuration has three fundamental concepts:

- **Templates** define the representation of a resource common across clusters
- **Placement** defines which clusters the resource is intended to appear in
- **Overrides** define per-cluster field-level variation to apply to the template

These three abstractions provide a concise representation of a resource intended
to appear in multiple clusters. They encode the minimum information required for
**propagation** and are well-suited to serve as the glue between any given
propagation mechanism and higher-order behaviors like policy-based placement and
dynamic scheduling.

These fundamental concepts provide building blocks that can be used by
higher-level APIs:

- **Status** collects the status of resources distributed by KubeFed across all federated clusters
- **Policy** determines which subset of clusters a resource is allowed to be distributed to
- **Scheduling** refers to a decision-making capability that can decide how
  workloads should be spread across different clusters similar to how a human
  operator would

## Features

| Feature                                                                                                         | Maturity | Feature Gate | Default |
|-----------------------------------------------------------------------------------------------------------------|----------|--------------|---------|
| [Push propagation of arbitrary types to remote clusters](./docs/userguide.md#verify-your-deployment-is-working) | Alpha | PushReconciler | true |
| [CLI utility (`kubefedctl`)](./docs/userguide.md#kubefedctl-cli)                                                | Alpha | | |
| [Generate KubeFed APIs without writing code](./docs/userguide.md#enabling-federation-of-an-api-type)            | Alpha | | |
| [Replica Scheduling Preferences](./docs/userguide.md#replicaschedulingpreference)                               | Alpha | SchedulerPreferences | true |

## Guides

### Quickstart

1. Clone this repo:
   ```
   git clone https://github.com/kubernetes-sigs/kubefed.git
   ```
1. Start a [kind](https://kind.sigs.k8s.io/) cluster:
   ```
   kind create cluster
   ```
1. Deploy kubefed:
   ```
   make deploy.kind
   ```

You now have a Kubernetes cluster with kubefed up and running. The cluster has been joined to itself and you can test federation of resources like this:

1. Verify the `KubeFedCluster` exists and is ready:
   ```
   kubectl -n kube-federation-system get kubefedcluster
   ```
   **If you're on macOS** the cluster will not immediately show as ready. You need to [change the API endpoint's URL first](https://github.com/kubernetes-sigs/kubefed/blob/master/docs/cluster-registration.md#joining-kind-clusters-on-macos):
   ```
   ./scripts/fix-joined-kind-clusters.sh
   ```
1. Create a namespace to be federated:
   ```
   kubectl create ns federate-me
   ```
1. Tell kubefed to federate that namespace (and the resources in it):
   ```
   ./bin/kubefedctl federate ns federate-me
   ```
1. Create a `ConfigMap` to be federated:
   ```
   kubectl -n federate-me create cm my-cm
   ```
1. Tell kubefed to federate that `ConfigMap`:
   ```
   ./bin/kubefedctl -n federate-me federate configmap my-cm
   ```
1. Verify the `FederatedConfigMap` has been created and propagates properly:
   ```
   kubectl -n federate-me describe federatedconfigmap my-cm
   ```

### User Guide

Take a look at our [user guide](docs/userguide.md) if you are interested in
using KubeFed.

### Development Guide

Take a look at our [development guide](docs/development.md) if you are
interested in contributing.

## Community

Refer to the [contributing guidelines](./CONTRIBUTING.md) if you would like to contribute to KubeFed.

### Communication channels

KubeFed is sponsored by [SIG Multicluster](https://github.com/kubernetes/community/tree/master/sig-multicluster) and it uses the same communication channels as SIG multicluster.

* Slack channel: [#sig-multicluster](http://slack.k8s.io/#sig-multicluster)
* [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-multicluster)

## Code of Conduct

Participation in the Kubernetes community is governed by the
[Kubernetes Code of Conduct](./code-of-conduct.md).
