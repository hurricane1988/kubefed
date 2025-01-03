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

package kubefedctl

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kubefed/pkg/version"
)

var (
	versionLong = `
		Version prints the version info of this command.`
	versionExample = `
		# Print kubefed command version
		kubefed version`
)

// NewCmdVersion prints out the release version info for this command binary.
func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the version info",
		Long:    versionLong,
		Example: versionExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(out, "kubefedctl version: %s\n", fmt.Sprintf("%#v", version.Get()))
		},
	}

	return cmd
}
