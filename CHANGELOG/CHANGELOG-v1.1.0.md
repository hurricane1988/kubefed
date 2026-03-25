# 📝 CHANGELOG - `kubefed`

## 📖 版本说明
---
## 🔧 优化
- 优化`version.go`获取`runtime` CPU核数方法为`runtime.GOMAXPROCS(0)`
- `go.opentelemetry.io/otel` 组件版本 `v1.37.0 -> v1.42.0`
- `go.opentelemetry.io/otel/metric` 组件版本 `v1.37.0 -> v1.42.0`
- `go.opentelemetry.io/otel/sdk` 组件版本 `v1.37.0 -> v1.42.0`
- `go.opentelemetry.io/otel/trace` 组件版本 `v1.37.0 -> v1.42.0`
- `go.opentelemetry.io/otel/trace` 组件版本 `v1.37.0 -> v1.42.0`
- `go.opentelemetry.io/proto/otlp` 组件版本 `v1.4.0 -> v1.10.0`
- `github.com/fatih/color` 组件版本 `v1.18.0 -> v1.19.0`
- `golang.org/x/net` 组件版本 `v0.50.0 -> v0.52.0`
- `go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc` 组件版本 `v0.58.0 -> v0.67.0`
- `go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp` 组件版本 `v0.58.0 -> v0.67.0`
- `go.opentelemetry.io/otel/exporters/otlp/otlptrace` 组件版本 `v1.33.0 -> v1.42.0`
- `google.golang.org/grpc` 组件版本 `v1.79.2 -> v1.79.3`
- `github.com/spf13/cobra` 组件版本 `v1.9.1 -> v1.10.2`
- `k8s.io/utils` 组件版本 `v0.0.0-20250604170112-4c0f3b243397 -> v0.0.0-20260319190234-28399d86e0b5`
- `google.golang.org/genproto/googleapis/rpc` 组件版本 `v0.0.0-20260209200024-4cfbd4190f57 -> v0.0.0-20260319201613-d00831a3d3e7`
- `golang.org/x/time` 组件版本 `v0.9.0 -> v0.15.0`

---
## ✨ 新增功能
`cmd/controller-manager/app/options/options.go` 配置知悉Election选举相关配置参数
- [x] Leader 选举参数绑定 `fs.DurationVar(&o.LeaderElection.LeaseDuration, "leader-elect-lease-duration", 15*time.Second,
		"The maximum duration that a leader can be stopped before it is replaced by another candidate.")`
- [x] `fs.DurationVar(&o.LeaderElection.RenewDeadline, "leader-elect-renew-deadline", 10*time.Second,
		"The interval between attempts by the acting master to renew a leadership slot before it stops leading.")`
- [x] `fs.DurationVar(&o.LeaderElection.RetryPeriod, "leader-elect-retry-period", 5*time.Second,
		"The duration the clients should wait between attempting acquisition and renewal of a leadership.")`
- [x] 绑定资源锁类型 (例如: "leases", "configmaps", "endpoints") `fs.StringVar((*string)(&o.LeaderElection.ResourceLock), "leader-elect-resource-lock", fedv1b1.LeasesResourceLock,
		"The type of resource object that will be used to lock during leader election.")`
- [x] 其他配置 `if o.ClusterHealthCheckConfig != nil {
		fs.DurationVar(&o.ClusterHealthCheckConfig.Period, "cluster-health-check-period", 10*time.Second, "How often to check the health of cluster.")
	}`

---
## ✅ 支持的 Kubernetes 版本

- `v1.22.x`
- `v1.23.x`
- `v1.24.x`
- `v1.25.x`
- `v1.26.x`
- `v1.27.x`
- `v1.28.x`
- `v1.29.x`
- `v1.30.x`
- `v1.31.x`
- `v1.32.x`
- `v1.33.x`
- `v1.35.x`

---
## 📦 构建与依赖版本
- 使用 Go **1.26.1**
- 设置 `GODEBUG`：`default=go1.26.1`

---
### 🔗 核心依赖

| **序号** | **组件库**                                                                     | **版本**                             |
|--------|-----------------------------------------------------------------------------|------------------------------------|
| 01     | k8s.io/apiextensions-apiserver                                              | v0.35.3                            |
| 02     | go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc | v0.58.0                            |
| 03     | k8s.io/api                                                                  | v0.35.3                            |
| 04     | sigs.k8s.io/controller-runtime                                              | v0.21.0                            |
| 05     | sigs.k8s.io/yaml                                                            | v1.6.0                             |
| 06     | sigs.k8s.io/structured-merge-diff/v4                                        | v4.6.0                             |
| 07     | k8s.io/client-go                                                            | v0.35.3                            |
| 08     | k8s.io/component-base                                                       | v0.35.3                            |
| 09     | k8s.io/apiserver                                                            | v0.35.3                            |
| 10     | k8s.io/apimachinery                                                         | v0.35.3                            |
| 11     | k8s.io/utils                                                                | v0.0.0-20260319190234-28399d86e0b5 |
| 12     | k8s.io/kubectl                                                              | v0.35.3                            |
| 13     | k8s.io/klog/v2                                                              | v2.140.0                           |
| 14     | sigs.k8s.io/controller-runtime                                              | v0.23.3                            |
| 15     | k8s.io/kube-openapi                                                         | v0.0.0-20260319004828-5883c5ee87b9 |
| 16     | sigs.k8s.io/structured-merge-diff/v6                                        | v6.3.2                             |
---
## 命令执行帮助信息
```shell

╭━╮╭━╮╱╱╭╮╭╮╱╱╱╱╭━━━┳╮╱╱╱╱╱╱╭╮╱╱╱╱╱╱╱╭╮╭━╮╱╱╭╮╱╱╱╱╱╭━╮╱╱╱╱╭╮
┃┃╰╯┃┃╱╱┃┣╯╰╮╱╱╱┃╭━╮┃┃╱╱╱╱╱╭╯╰╮╱╱╱╱╱╱┃┃┃╭╯╱╱┃┃╱╱╱╱╱┃╭╯╱╱╱╱┃┃
┃╭╮╭╮┣╮╭┫┣╮╭╋╮╱╱┃┃╱╰┫┃╭╮╭┳━┻╮╭╋━━┳━╮╱┃╰╯╯╭╮╭┫╰━┳━━┳╯╰┳━━┳━╯┃
┃┃┃┃┃┃┃┃┃┃┃┃┣╋━━┫┃╱╭┫┃┃┃┃┃━━┫┃┃┃━┫╭┻━┫╭╮┃┃┃┃┃╭╮┃┃━╋╮╭┫┃━┫╭╮┃
┃┃┃┃┃┃╰╯┃╰┫╰┫┣━━┫╰━╯┃╰┫╰╯┣━━┃╰┫┃━┫┣━━┫┃┃╰┫╰╯┃╰╯┃┃━┫┃┃┃┃━┫╰╯┃
╰╯╰╯╰┻━━┻━┻━┻╯╱╱╰━━━┻━┻━━┻━━┻━┻━━┻╯╱╱╰╯╰━┻━━┻━━┻━━╯╰╯╰━━┻━━╯

+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+-------------+---------------+--------------+
| COMMUNITY  | AUTHOR       | VERSION | GIT COMMIT                               | BUILD DATE| GO VERSION | COMPILER | PLATFORM    | RUNTIME CORES | TOTAL MEMORY |
+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+-------------+---------------+--------------+
| CodeFuture | Jianping Niu | v1.1.0  | 28771945f64eaa8c65c136c59283b2e02350b447 | 2026-03-25T01:48:34Z| go1.26.1   | gc       | linux/amd64 | 1 cores       | 8919 KB      |
+------------+--------------+---------+------------------------------------------+----------------------+------------+----------+-------------+---------------+--------------+
The KubeFed controller manager runs a bunch of controllers
which watch KubeFed CRDs and the corresponding resources in
member clusters and do the necessary reconciliation

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
      --leader-elect-lease-duration duration   The maximum duration that a leader can be stopped beforeit is replaced by another candidate. (default 15s)
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
      --rest-config-qps float32                Maximum QPS to the api-server from this client. (default100)
      --skip_headers                           If true, avoid header prefixes in the log messages
      --skip_log_headers                       If true, avoid headers when opening log files (no effectwhen -logtostderr=true)
      --stderrthreshold severity               logs at or above this threshold go to stderr when writing to files and stderr (no effect when -logtostderr=true or -alsologtostderr=true unless -legacy_stderr_threshold_behavior=false) (default 2)
  -v, --v Level                                number for the log level verbosity
      --version                                Prints the Version info of controller-manager.
      --vmodule moduleSpec                     comma-separated list of pattern=N settings for file-filtered logging
```

---
## 🧱项目工程结构
```shell
.
├── CHANGELOG
│   ├── CHANGELOG-v1.0.1.md
│   ├── CHANGELOG-v1.0.2.md
│   └── CHANGELOG-v1.1.0.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── OWNERS
├── PROJECT
├── README.md
├── SECURITY.md
├── SECURITY_CONTACTS
├── bin
│   ├── controller-manager
│   ├── hyperfed
│   ├── kubefedctl
│   └── webhook
├── build
│   └── kubefed
│       └── Dockerfile
├── charts
│   ├── index.yaml
│   ├── kubefed
│   │   ├── Chart.yaml
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── charts
│   │   │   └── controllermanager
│   │   ├── crds
│   │   │   └── crds.yaml
│   │   ├── templates
│   │   │   ├── _helpers.tpl
│   │   │   └── federatedtypeconfig.yaml
│   │   └── values.yaml
│   └── service-monitor.yaml
├── cmd
│   ├── controller-manager
│   │   ├── app
│   │   │   ├── controller-manager.go
│   │   │   ├── leaderelection
│   │   │   └── options
│   │   └── main.go
│   ├── hyperfed
│   │   └── main.go
│   ├── kubefedctl
│   │   └── main.go
│   └── webhook
│       ├── app
│       │   └── webhook.go
│       └── main.go
├── code-of-conduct.md
├── config
│   ├── crds
│   │   ├── core.kubefed.io_clusterpropagatedversions.yaml
│   │   ├── core.kubefed.io_federatedservicestatuses.yaml
│   │   ├── core.kubefed.io_federatedtypeconfigs.yaml
│   │   ├── core.kubefed.io_kubefedclusters.yaml
│   │   ├── core.kubefed.io_kubefedconfigs.yaml
│   │   ├── core.kubefed.io_propagatedversions.yaml
│   │   └── scheduling.kubefed.io_replicaschedulingpreferences.yaml
│   ├── enabletypedirectives
│   │   ├── clusterroles.rbac.authorization.k8s.io.yaml
│   │   ├── configmaps.yaml
│   │   ├── deployments.apps.yaml
│   │   ├── ingresses.networking.k8s.io.yaml
│   │   ├── jobs.batch.yaml
│   │   ├── namespaces.yaml
│   │   ├── replicasets.apps.yaml
│   │   ├── secrets.yaml
│   │   ├── serviceaccounts.yaml
│   │   └── services.yaml
│   └── kubefedconfig.yaml
├── docs
│   ├── OWNERS
│   ├── cluster-registration.md
│   ├── concepts.md
│   ├── development.md
│   ├── environments
│   │   ├── gke.md
│   │   ├── icp.md
│   │   ├── kind.md
│   │   └── minikube.md
│   ├── images
│   │   ├── concepts.png
│   │   ├── ingressdns-with-externaldns.png
│   │   ├── propagation.png
│   │   └── servicedns-with-externaldns.png
│   ├── installation.md
│   ├── keps
│   │   ├── 20200302-kubefed-metrics.md
│   │   ├── 20200619-federated-resource-status.md
│   │   ├── 20200619-kubefed-pull-reconciliation.md
│   │   └── images
│   │       ├── kubefedArch.jpg
│   │       ├── kubefedv2Example.jpg
│   │       └── kubefedv2_seconditeration.jpg
│   ├── project-layout.md
│   ├── releasing.md
│   └── userguide.md
├── example
│   ├── config
│   │   └── kubefedconfig.yaml
│   ├── sample
│   │   ├── configmap.yaml
│   │   ├── deployment.yaml
│   │   ├── federatedclusterrole.yaml
│   │   ├── federatedclusterrolebinding.yaml
│   │   ├── federatedconfigmap.yaml
│   │   ├── federateddeployment.yaml
│   │   ├── federatedingress.yaml
│   │   ├── federatedjob.yaml
│   │   ├── federatednamespace.yaml
│   │   ├── federatedsecret.yaml
│   │   ├── federatedservice.yaml
│   │   ├── federatedserviceaccount.yaml
│   │   ├── namespace.yaml
│   │   └── service.yaml
│   └── scheduling
│       └── federateddeployment-rsp.yaml
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
│   │   ├── addtoscheme_core_v1alpha1.go
│   │   ├── addtoscheme_core_v1beta1.go
│   │   ├── addtoscheme_scheduling_v1alpha1.go
│   │   ├── apis.go
│   │   ├── core
│   │   │   ├── common
│   │   │   ├── group.go
│   │   │   ├── typeconfig
│   │   │   ├── v1alpha1
│   │   │   └── v1beta1
│   │   └── scheduling
│   │       ├── group.go
│   │       └── v1alpha1
│   ├── client
│   │   └── generic
│   │       ├── genericclient.go
│   │       └── scheme
│   ├── constants
│   │   └── constants.go
│   ├── controller
│   │   ├── doc.go
│   │   ├── federatedtypeconfig
│   │   │   └── controller.go
│   │   ├── kubefedcluster
│   │   │   ├── clusterclient.go
│   │   │   ├── controller.go
│   │   │   ├── controller_integration_test.go
│   │   │   └── controller_test.go
│   │   ├── schedulingmanager
│   │   │   └── controller.go
│   │   ├── schedulingpreference
│   │   │   └── controller.go
│   │   ├── status
│   │   │   └── controller.go
│   │   ├── sync
│   │   │   ├── accessor.go
│   │   │   ├── controller.go
│   │   │   ├── dispatch
│   │   │   ├── resource.go
│   │   │   ├── resource_test.go
│   │   │   ├── status
│   │   │   └── version
│   │   ├── testdata
│   │   │   └── fixtures
│   │   ├── utils
│   │   │   ├── backoff.go
│   │   │   ├── cluster_util.go
│   │   │   ├── constants.go
│   │   │   ├── controllerconfig.go
│   │   │   ├── delaying_deliverer.go
│   │   │   ├── delaying_deliverer_test.go
│   │   │   ├── deletionannotation.go
│   │   │   ├── deletionannotation_test.go
│   │   │   ├── federated_informer.go
│   │   │   ├── federatedstatus.go
│   │   │   ├── finalizers
│   │   │   ├── genericinformer.go
│   │   │   ├── handlers.go
│   │   │   ├── handlers_test.go
│   │   │   ├── managedlabel.go
│   │   │   ├── meta.go
│   │   │   ├── meta_test.go
│   │   │   ├── naming.go
│   │   │   ├── orphaninganotation.go
│   │   │   ├── overrides.go
│   │   │   ├── placement.go
│   │   │   ├── placement_test.go
│   │   │   ├── planner
│   │   │   ├── podanalyzer
│   │   │   ├── propagatedversion.go
│   │   │   ├── qualifiedname.go
│   │   │   ├── resourceclient.go
│   │   │   ├── resourceinformer.go
│   │   │   ├── safe_map.go
│   │   │   ├── worker.go
│   │   │   └── worker_test.go
│   │   └── webhook
│   │       ├── federatedtypeconfig
│   │       ├── kubefedcluster
│   │       ├── kubefedconfig
│   │       └── util.go
│   ├── doc.go
│   ├── features
│   │   └── features.go
│   ├── kubefedctl
│   │   ├── disable.go
│   │   ├── enable
│   │   │   ├── deprecatedapis.go
│   │   │   ├── deprecatedapis_test.go
│   │   │   ├── directive.go
│   │   │   ├── enable.go
│   │   │   ├── schema.go
│   │   │   ├── util.go
│   │   │   └── validation.go
│   │   ├── federate
│   │   │   ├── federate.go
│   │   │   ├── federate_test.go
│   │   │   └── util.go
│   │   ├── join.go
│   │   ├── join_test.go
│   │   ├── kubefedctl.go
│   │   ├── options
│   │   │   └── options.go
│   │   ├── orphaning
│   │   │   ├── disable.go
│   │   │   ├── enable.go
│   │   │   ├── orphaning.go
│   │   │   └── status.go
│   │   ├── suite_test.go
│   │   ├── unjoin.go
│   │   ├── util
│   │   │   ├── util.go
│   │   │   └── yaml_writer.go
│   │   └── version.go
│   ├── metrics
│   │   └── metrics.go
│   ├── schedulingtypes
│   │   ├── interface.go
│   │   ├── plugin.go
│   │   ├── plugin_test.go
│   │   ├── replicascheduler.go
│   │   ├── resources.go
│   │   └── typeregistry.go
│   └── version
│       └── version.go
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
│   │   ├── bindata.go
│   │   ├── crudtester.go
│   │   ├── fixtures
│   │   │   ├── clusterroles.rbac.authorization.k8s.io.yaml
│   │   │   ├── configmaps.yaml
│   │   │   ├── deployments.apps.yaml
│   │   │   ├── ingresses.networking.k8s.io.yaml
│   │   │   ├── jobs.batch.yaml
│   │   │   ├── namespaces.yaml
│   │   │   ├── replicasets.apps.yaml
│   │   │   ├── secrets.yaml
│   │   │   ├── serviceaccounts.yaml
│   │   │   └── services.yaml
│   │   ├── fixtures.go
│   │   ├── resource_helper.go
│   │   ├── testobjects.go
│   │   ├── typeconfig.go
│   │   ├── util.go
│   │   └── validation.go
│   └── e2e
│       ├── crd.go
│       ├── crud.go
│       ├── defaulting.go
│       ├── deleteoptions.go
│       ├── e2e.go
│       ├── e2e_test.go
│       ├── federate.go
│       ├── framework
│       │   ├── cleanup.go
│       │   ├── cluster.go
│       │   ├── controller.go
│       │   ├── enable.go
│       │   ├── framework.go
│       │   ├── ginkgowrapper
│       │   ├── logger.go
│       │   ├── test_context.go
│       │   ├── unmanaged.go
│       │   ├── util.go
│       │   └── wait.go
│       ├── ftccontroller.go
│       ├── kubefedcluster.go
│       ├── leaderelection.go
│       ├── not_ready.go
│       ├── placement.go
│       ├── scale.go
│       ├── schedulermanager.go
│       ├── scheduling.go
│       ├── validation.go
│       └── version.go
├── third-party
│   └── k8s.io
│       └── repo-infra
│           ├── BUILD.bazel
│           ├── CONTRIBUTING.md
│           ├── LICENSE
│           ├── OWNERS
│           ├── README.md
│           ├── SECURITY.md
│           ├── SECURITY_CONTACTS
│           ├── WORKSPACE
│           ├── cmd
│           ├── code-of-conduct.md
│           ├── defs
│           ├── go
│           ├── go.mod
│           ├── go.sum
│           ├── hack
│           ├── load.bzl
│           ├── presubmit.sh
│           ├── repos.bzl
│           ├── tools
│           └── verify
└── tools
    ├── go.mod
    ├── go.sum
    └── tools.go
```
