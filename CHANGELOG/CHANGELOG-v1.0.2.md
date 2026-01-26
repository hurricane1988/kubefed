# 📝 CHANGELOG - `kubefed`

## 📖 版本说明
---
## 🔧 优化
- 优化`version.go`获取`runtime` CPU核数方法为`runtime.GOMAXPROCS(0)`
- `cmd` 启用执行 `main.go` 均添加`go.uber.org/automaxprocs/maxprocs`,以便准确获取当前容器实际CPU核数
- `golang.org/x/net` 组件版本 `v0.41.0 -> v0.48.0`
- `golang.org/x/text` 组件版本 `v0.26.0 -> v0.33.0`
- `golang.org/x/sync` 组件版本 `v0.15.0 -> v0.19.0`
- `golang.org/x/sys` 组件版本 `v0.33.0 -> v0.40.0`
- `golang.org/x/term` 组件版本 `v0.32.0 -> v0.39.0`

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

---
## 📦 构建与依赖版本
- 使用 Go **1.25.5**
- 设置 `GODEBUG`：`default=go1.25.5`

---
### 🔗 核心依赖

| **序号** | **组件库**                                                                     | **版本**                             |
|--------|-----------------------------------------------------------------------------|------------------------------------|
| 01     | k8s.io/apiextensions-apiserver                                              | v0.33.0                            |
| 02     | go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc | v0.58.0                            |
| 03     | k8s.io/api                                                                  | v0.33.0                            |
| 04     | sigs.k8s.io/controller-runtime                                              | v0.21.0                            |
| 05     | sigs.k8s.io/yaml                                                            | v1.6.0                             |
| 06     | sigs.k8s.io/structured-merge-diff/v4                                        | v4.6.0                             |
| 07     | k8s.io/client-go                                                            | v0.33.0                            |
| 08     | k8s.io/component-base                                                       | v0.33.0                            |
| 09     | k8s.io/apiserver                                                            | v0.33.0                            |
| 10     | k8s.io/apimachinery                                                         | v0.33.0                            |
| 11     | k8s.io/utils                                                                | v0.0.0-20250604170112-4c0f3b243397 |
| 12     | k8s.io/kubectl                                                              | v0.33.0                            |

---
## 🧱项目工程结构
```shell
.
├── CHANGELOG
│   ├── CHANGELOG-v1.0.1.md
│   └── CHANGELOG-v1.0.2.md
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
│   ├── sample1
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
