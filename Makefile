# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Image URL to use all building/pushing image targets
VERSION ?= v1.0.1
IMG ?= codefuthure/kubefed

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
# Be aware that the target commands are only tested with Docker which is
# scaffolded by default. However, you might want to replace it to use other
# tools. (i.e. podman)
CONTAINER_TOOL ?= docker

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_TREE_STATE ?= $(shell if git diff --quiet && git diff --cached --quiet; then echo clean; else echo dirty; fi)

LDFLAG_OPTIONS = -ldflags "-X sigs.k8s.io/kubefed/pkg/version.Version=$(VERSION) \
                      -X sigs.k8s.io/kubefed/pkg/version.GitCommit=$(GIT_COMMIT) \
                      -X sigs.k8s.io/kubefed/pkg/version.GitTreeState=$(GIT_TREE_STATE) \
                      -X sigs.k8s.io/kubefed/pkg/version.BuildDate=$(BUILDDATE)"

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Versions
KUSTOMIZE_VERSION ?= v5.3.0
CONTROLLER_TOOLS_VERSION ?= v0.14.0
ENVTEST_VERSION ?= latest
GOLANGCI_LINT_VERSION ?= v1.54.2

## Tool Binaries
KUBECTL ?= kubectl
KUSTOMIZE ?= $(LOCALBIN)/kustomize-$(KUSTOMIZE_VERSION)
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen-$(CONTROLLER_TOOLS_VERSION)
ENVTEST ?= $(LOCALBIN)/setup-envtest-$(ENVTEST_VERSION)
GOLANGCI_LINT = $(LOCALBIN)/golangci-lint-$(GOLANGCI_LINT_VERSION)


# Retrieve the Git version information, including any local changes
GIT_VERSION ?= $(shell git describe --always --dirty)
# Retrieve the exact Git tag if it matches the current commit, otherwise empty
GIT_TAG ?= $(shell git describe --tags --exact-match 2>/dev/null)
# Retrieve the full Git commit hash
GIT_HASH ?= $(shell git rev-parse HEAD)
# Retrieve the current Git branch name, excluding 'HEAD' if present
GIT_BRANCH ?= $(filter-out HEAD,$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null))
# Get the build date in UTC format
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: all

all: help

##@ Development
ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

.PHONY: manifests
manifests: generate-code ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) crd webhook paths="./pkg/apis/..." output:crd:artifacts:config=config/crds

.PHONY: generate-code
generate-code: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

.PHONY: generate
generate: manifests build ## Generate the kubefedctl and generate-code.
	./scripts/sync-up-helm-chart.sh
	./scripts/update-bindata.sh

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: ## Run unit test.
	go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
	source <(setup-envtest use -p env 1.24.x) && \
		go test $(TEST_PKGS)

##@ Build

.PHONY: build
build: fmt vet ## build hyperfed controller kubefedctl webhook binary.
	go build $(LDFLAG_OPTIONS) -o bin/controller-manager cmd/controller-manager/main.go
	go build $(LDFLAG_OPTIONS) -o bin/hyperfed cmd/hyperfed/main.go
	go build $(LDFLAG_OPTIONS) -o bin/kubefedctl cmd/kubefedctl/main.go
	go build $(LDFLAG_OPTIONS) -o bin/webhook cmd/webhook/main.go

.PHONY: lint
lint: ## Run golangci-lint check.
	golangci-lint run -c .golangci.yml --fix

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the kubefed.
	$(CONTAINER_TOOL) build -t ${IMG}:${VERSION} -f build/kubefed/Dockerfile \
							--build-arg VERSION=$(VERSION) \
							--build-arg GITCOMMIT=$(GIT_COMMIT) \
							--build-arg GITTREESTATE=$(GIT_TREE_STATE) \
							--build-arg BUILDDATE=$(BUILDDATE) .

.PHONY: docker-push
docker-push: ## Push docker image with the kubefed.
	$(CONTAINER_TOOL) ${IMG}:${VERSION}

# PLATFORMS defines the target platforms for the manager image be built to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - be able to use docker buildx. More info: https://docs.docker.com/build/buildx/
# - have enabled BuildKit. More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image to your registry (i.e. if you do not set a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To adequately provide solutions that are compatible with multiple platforms, you should consider using this option.
PLATFORMS ?= linux/amd64,linux/arm64

.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the kubefed for cross-platform support.
	- $(CONTAINER_TOOL) buildx create --name kubefed
	$(CONTAINER_TOOL) buildx use kubefed
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --tag ${IMG}:${VERSION} -f build/kubefed/Dockerfile \
 											--build-arg VERSION=$(VERSION) --build-arg GITCOMMIT=$(GIT_COMMIT) --build-arg GIT_TREE_STATE=$(GIT_TREE_STATE) --build-arg BUILDDATE=$(BUILDDATE) .
	- $(CONTAINER_TOOL) buildx rm kubefed


##@ Dependencies

.PHONY: envtest
envtest: $(ENVTEST) ## Download setup-envtest locally if necessary.
$(ENVTEST): $(LOCALBIN)
	$(call go-install-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest,$(ENVTEST_VERSION))

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(GOLANGCI_LINT): $(LOCALBIN)
	$(call go-install-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint,${GOLANGCI_LINT_VERSION})

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef

# Generate code

.PHONY: clean
clean: ## Clean all the binaries.
	@if [ -d "bin" ];then rm -rf bin; fi
	$(CONTAINER_TOOL) rmi ${IMG}

.PHONY: deploy.kind
deploy.kind: generate ## Deploy the kubefed.
	KIND_LOAD_IMAGE=y FORCE_REDEPLOY=y ./scripts/deploy-kubefed.sh $(IMAGE_NAME)
