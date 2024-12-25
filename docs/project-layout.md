```shell
.
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── OWNERS
├── PROJECT
├── README.md
├── SECURITY.md
├── SECURITY_CONTACTS
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
│   │   │       ├── Chart.yaml
│   │   │       ├── crds
│   │   │       └── templates
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
│   │   │   │   └── leaderelection.go
│   │   │   └── options
│   │   │       └── options.go
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
├── controller-manager.log
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
│   │   │   │   ├── constants.go
│   │   │   │   └── utils.go
│   │   │   ├── group.go
│   │   │   ├── typeconfig
│   │   │   │   ├── interface.go
│   │   │   │   └── utils.go
│   │   │   ├── v1alpha1
│   │   │   │   ├── clusterpropagatedversion_types.go
│   │   │   │   ├── federatedservicestatus_types.go
│   │   │   │   ├── groupversion_info.go
│   │   │   │   ├── propagatedversion_types.go
│   │   │   │   └── zz_generated.deepcopy.go
│   │   │   └── v1beta1
│   │   │       ├── defaults
│   │   │       ├── federatedtypeconfig_types.go
│   │   │       ├── groupversion_info.go
│   │   │       ├── kubefedcluster_types.go
│   │   │       ├── kubefedcluster_types_test.go
│   │   │       ├── kubefedconfig_types.go
│   │   │       ├── suite_test.go
│   │   │       ├── validation
│   │   │       └── zz_generated.deepcopy.go
│   │   └── scheduling
│   │       ├── group.go
│   │       └── v1alpha1
│   │           ├── groupversion_info.go
│   │           ├── replicaschedulingpreference_types.go
│   │           └── zz_generated.deepcopy.go
│   ├── client
│   │   └── generic
│   │       ├── genericclient.go
│   │       └── scheme
│   │           └── register.go
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
│   │   │   │   ├── checkunmanaged.go
│   │   │   │   ├── managed.go
│   │   │   │   ├── operation.go
│   │   │   │   ├── retain.go
│   │   │   │   ├── retain_test.go
│   │   │   │   └── unmanaged.go
│   │   │   ├── resource.go
│   │   │   ├── resource_test.go
│   │   │   ├── status
│   │   │   │   ├── status.go
│   │   │   │   └── status_test.go
│   │   │   └── version
│   │   │       ├── adapter.go
│   │   │       ├── cluster.go
│   │   │       ├── manager.go
│   │   │       └── namespaced.go
│   │   ├── testdata
│   │   │   └── fixtures
│   │   │       └── crds.yaml
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
│   │   │   │   ├── finalizers.go
│   │   │   │   └── finalizers_test.go
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
│   │   │   │   ├── planner.go
│   │   │   │   └── planner_test.go
│   │   │   ├── podanalyzer
│   │   │   │   ├── pod_helper.go
│   │   │   │   └── pod_helper_test.go
│   │   │   ├── propagatedversion.go
│   │   │   ├── qualifiedname.go
│   │   │   ├── resourceclient.go
│   │   │   ├── resourceinformer.go
│   │   │   ├── safe_map.go
│   │   │   ├── worker.go
│   │   │   └── worker_test.go
│   │   └── webhook
│   │       ├── federatedtypeconfig
│   │       │   └── webhook.go
│   │       ├── kubefedcluster
│   │       │   └── webhook.go
│   │       ├── kubefedconfig
│   │       │   └── webhook.go
│   │       └── utils.go
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
│   │   │   ├── utils.go
│   │   │   └── validation.go
│   │   ├── federate
│   │   │   ├── federate.go
│   │   │   ├── federate_test.go
│   │   │   └── utils.go
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
│   │   ├── utils
│   │   │   ├── utils.go
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
│       ├── base.go
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
│   └── utils.sh
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
│   │   ├── utils.go
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
│       │   │   └── wrapper.go
│       │   ├── logger.go
│       │   ├── test_context.go
│       │   ├── unmanaged.go
│       │   ├── utils.go
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
│           │   └── kazel
│           ├── code-of-conduct.md
│           ├── defs
│           │   ├── BUILD.bazel
│           │   ├── build.bzl
│           │   ├── deb.bzl
│           │   ├── diff_test.sh
│           │   ├── gcs_uploader.py
│           │   ├── go.bzl
│           │   ├── pkg.bzl
│           │   ├── rpm.bzl
│           │   ├── run_in_workspace.bzl
│           │   ├── testdata
│           │   ├── testgen
│           │   └── testpkg
│           ├── go
│           │   ├── BUILD.bazel
│           │   └── sdk_versions.bzl
│           ├── go.mod
│           ├── go.sum
│           ├── hack
│           │   ├── BUILD.bazel
│           │   ├── tools.go
│           │   ├── update-bazel.sh
│           │   ├── update-deps.sh
│           │   ├── update-gofmt.sh
│           │   ├── verify-bazel.sh
│           │   ├── verify-deps.sh
│           │   ├── verify-gofmt.sh
│           │   ├── verify-golangci-lint.sh
│           │   ├── verify_boilerplate.py
│           │   └── verify_boilerplate_test.py
│           ├── load.bzl
│           ├── presubmit.sh
│           ├── repos.bzl
│           ├── tools
│           │   ├── CROSSTOOL
│           │   └── build_tar
│           └── verify
│               ├── BUILD.bazel
│               ├── README.md
│               ├── boilerplate
│               ├── go_install_from_commit.sh
│               ├── verify-bazel.sh
│               ├── verify-boilerplate.sh
│               ├── verify-errexit.sh
│               └── verify-go-src.sh
└── tools
    ├── go.mod
    ├── go.sum
    └── tools.go
```
