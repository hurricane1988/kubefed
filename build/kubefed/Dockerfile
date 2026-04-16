# Copyright 2026 The CodeFuture Authors.
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
ARG GO_VERSION=1.26.2-alpine3.23
FROM golang:${GO_VERSION} AS builder

# Set Go Proxy for faster dependency downloads
ARG GOPROXY="https://goproxy.cn,direct"
ENV GOPROXY=${GOPROXY}

WORKDIR /workspace

# Copy dependency files first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build arguments passed by the build system
ARG TARGETOS
ARG VERSION
ARG GITCOMMIT
ARG GITTREESTATE
ARG BUILDDATE
ARG TARGETARCH

# Build the binary with optimization flags
# -s -w: Omit symbol table and debug information to reduce binary size
# -trimpath: Remove file system paths from the resulting binary for better reproducibility
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags "-s -w \
              -X sigs.k8s.io/kubefed/pkg/version.Version=${VERSION} \
              -X sigs.k8s.io/kubefed/pkg/version.GitCommit=${GITCOMMIT} \
              -X sigs.k8s.io/kubefed/pkg/version.GitTreeState=${GITTREESTATE} \
              -X sigs.k8s.io/kubefed/pkg/version.BuildDate=${BUILDDATE}" \
    -a -o hyperfed cmd/hyperfed/main.go

# Create symbolic links in the builder stage because the runtime image (distroless)
# does not contain a shell or the 'ln' command.
RUN ln -s hyperfed controller-manager \
 && ln -s hyperfed kubefedctl \
 && ln -s hyperfed webhook

# --- Stage 2: Prepare the runtime environment ---
# Distroless static images are highly secure as they contain only the application and its dependencies
FROM gcr.io/distroless/static-debian13:latest
# Set the working directory
WORKDIR /hyperfed/

# Copy the binary built in the first stage
COPY --from=builder /workspace/hyperfed /hyperfed/
COPY --from=builder /workspace/controller-manager /hyperfed/
COPY --from=builder /workspace/kubefedctl /hyperfed/
COPY --from=builder /workspace/webhook /hyperfed/

# Use UID 65534 (nobody) for security compliance in restricted Kubernetes environments
USER 65534:65534

# Set the default entrypoint.
# Since distroless has no shell, we must use the exec form and point to the binary directly.
ENTRYPOINT ["/hyperfed/hyperfed"]
