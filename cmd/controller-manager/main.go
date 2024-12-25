/*
Copyright 2024 The CodeFuture Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"sigs.k8s.io/kubefed/pkg/version"

	_ "sigs.k8s.io/controller-runtime/pkg/metrics" // for work queue metrics registration

	genericapiserver "k8s.io/apiserver/pkg/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Load all client auth plugins for GCP, Azure, Openstack, etc
	"k8s.io/component-base/logs"

	"sigs.k8s.io/kubefed/cmd/controller-manager/app"
)

// Controller-manager main.
func main() {
	// Print the terminal information.
	fmt.Println(version.Term())
	// The core code for initializing logs using logs.InitLogs() sets up the configuration for klog by invoking
	// klog.InitFlags(nil) or a similar initialization method.
	// It relies on the functionality provided by the klog library at its core.
	logs.InitLogs()
	// It is used to ensure that all log data is written to the target output (such as files or standard output)
	// when the program exits.It is typically used in conjunction with InitLogs() to prevent log loss.
	defer logs.FlushLogs()

	stopChan := genericapiserver.SetupSignalHandler()

	if err := app.NewControllerManagerCommand(stopChan).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
