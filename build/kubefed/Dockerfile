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
FROM golang:1.23.2 AS builder
# Set the working directory in the builder stage
WORKDIR /build
# Copy the application source code into the build stage
ADD . /build/
ENV GOPROXY=https://goproxy.cn,direct
# Build the binary for the application
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -o /tmp/hyperfed cmd/hyperfed/main.go

# --- Stage 2: Prepare the runtime environment ---
FROM alpine:3.20.0
# Add a non-root user for security
RUN adduser -D -g hyperfed -u 1001 hyperfed
# Create the application directory
RUN mkdir -p /hyperfed
# Set the working directory
WORKDIR /hyperfed/
# Copy the binary built in the first stage
COPY --from=builder /tmp/ /hyperfed/

RUN ln -s hyperfed controller-manager \
 && ln -s hyperfed kubefedctl \
 && ln -s hyperfed webhook \

# Change ownership of the directory to the non-root user
RUN chown -R hyperfed:hyperfed /hyperfed
# Switch to the non-root user
USER hyperfed
# Set the default entrypoint for the container
CMD ["sh"]
