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

# --- Stage 1: Build the application ---
ARG GO_VERSION=1.24.4-alpine3.22
FROM golang:${GO_VERSION} AS builder

# Set the working directory in the builder stage
WORKDIR /build
# Copy the application source code into the build stage
COPY . .

# This is the architecture you're building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETOS
ARG VERSION
ARG GITCOMMIT
ARG GITTREESTATE
ARG BUILDDATE
ARG TARGETARCH
# NOTICE: ARG 变量中的带有逗号、分号或特殊字符的值加上引号
ARG GOPROXY="https://goproxy.cn,direct"

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build \
    -ldflags "-X sigs.k8s.io/kubefed/pkg/version.Version=${VERSION} \
              -X sigs.k8s.io/kubefed/pkg/version.GitCommit=${GITCOMMIT} \
              -X sigs.k8s.io/kubefed/pkg/version.GitTreeState=${GITTREESTATE} \
              -X sigs.k8s.io/kubefed/pkg/version.BuildDate=${BUILDDATE}" -a -o /tmp/hyperfed cmd/hyperfed/main.go

# --- Stage 2: Prepare the runtime environment ---
FROM alpine:3.22.1

# Install the curl tool
RUN apk add --no-cache curl

# Set the working directory
WORKDIR /hyperfed/

# Create a non-privileged user that the app will run under.
# See https://docs.docker.com/go/dockerfile-user-best-practices/
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    hyperfed

# Create the application directory
RUN mkdir -p /hyperfed

# Copy the binary built in the first stage
COPY --from=builder /tmp/ /hyperfed/

RUN ln -s hyperfed controller-manager \
 && ln -s hyperfed kubefedctl \
 && ln -s hyperfed webhook

# Setup the app user
USER hyperfed

# Set the default entrypoint for the container
CMD ["sh"]
