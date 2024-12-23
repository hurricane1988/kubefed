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
IMG ?= kubefed:latest

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
	go build -o bin/controller-manager cmd/controller-manager/main.go
	go build -o bin/hyperfed cmd/hyperfed/main.go
	go build -o bin/kubefedctl cmd/kubefedctl/main.go
	go build -o bin/webhook cmd/webhook/main.go

.PHONY: lint
lint: ## Run golangci-lint check.
	golangci-lint run -c .golangci.yml --fix

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the kubefed.
	$(CONTAINER_TOOL) build -t ${IMG} -f build/kubefed/Dockerfile .

.PHONY: docker-push
docker-push: ## Push docker image with the kubefed.
	$(CONTAINER_TOOL) push ${IMG}

# PLATFORMS defines the target platforms for the manager image be built to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - be able to use docker buildx. More info: https://docs.docker.com/build/buildx/
# - have enabled BuildKit. More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image to your registry (i.e. if you do not set a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To adequately provide solutions that are compatible with multiple platforms, you should consider using this option.
PLATFORMS ?= linux/amd64,linux/arm64

.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the kubefed for cross-platform support.
	- $(CONTAINER_TOOL) buildx create --name project-v3-builder
	$(CONTAINER_TOOL) buildx use project-v3-builder
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f build/kubefed/Dockerfile .
	- $(CONTAINER_TOOL) buildx rm project-v3-builder

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

# Generate code
.PHONY: generate-code
generate-code: controller-gen ## Run controller-gen and generate code.
	controller-gen object:headerFile=./hack/boilerplate.go.txt paths="./..."

# Generate code
.PHONY: generate
generate: generate-code build ## Generate the kubefedctl and generate-code.
	./scripts/sync-up-helm-chart.sh
	./scripts/update-bindata.sh

.PHONY: clean
clean: ## Clean all the binaries.
	@if [ -d "bin" ];then rm -rf bin; fi
	$(CONTAINER_TOOL) rmi ${IMG}

.PHONY: controller-gen
controller-gen: ## Install the controller-gen tool.
	command -v controller-gen &> /dev/null || (cd tools && go install sigs.k8s.io/controller-tools/cmd/controller-gen)

.PHONY: deploy.kind
deploy.kind: generate ## Deploy the kubefed.
	KIND_LOAD_IMAGE=y FORCE_REDEPLOY=y ./scripts/deploy-kubefed.sh $(IMAGE_NAME)
