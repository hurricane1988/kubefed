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

package version

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
	"sigs.k8s.io/kubefed/pkg/constants"
)

type Info struct {
	Version      string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
	RuntimeCores int    `json:"RuntimeCores"`
	TotalMem     int    `json:"TotalMem"`
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return Info{
		Version:      constants.DefaultVersion,
		GitCommit:    constants.DefaultGitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    constants.DefaultBuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		RuntimeCores: runtime.NumCPU(),
		TotalMem:     int(memStats.TotalAlloc / 1024),
	}
}

var (
	Yellow       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	YellowItalic = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintFunc()
	Green        = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	Blue         = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Cyan         = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	Red          = color.New(color.FgHiRed, color.Bold).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	WhiteBold    = color.New(color.FgWhite, color.Bold).SprintFunc()
	forceDetail  = "yaml"
)

// Term Print the terminal
func Term() string {
	return fmt.Sprint(Blue(`
╭━╮╭━╮╱╱╭╮╭╮╱╱╱╱╭━━━┳╮╱╱╱╱╱╱╭╮╱╱╱╱╱╱╱╭╮╭━╮╱╱╭╮╱╱╱╱╱╭━╮╱╱╱╱╭╮
┃┃╰╯┃┃╱╱┃┣╯╰╮╱╱╱┃╭━╮┃┃╱╱╱╱╱╭╯╰╮╱╱╱╱╱╱┃┃┃╭╯╱╱┃┃╱╱╱╱╱┃╭╯╱╱╱╱┃┃
┃╭╮╭╮┣╮╭┫┣╮╭╋╮╱╱┃┃╱╰┫┃╭╮╭┳━┻╮╭╋━━┳━╮╱┃╰╯╯╭╮╭┫╰━┳━━┳╯╰┳━━┳━╯┃
┃┃┃┃┃┃┃┃┃┃┃┃┣╋━━┫┃╱╭┫┃┃┃┃┃━━┫┃┃┃━┫╭┻━┫╭╮┃┃┃┃┃╭╮┃┃━╋╮╭┫┃━┫╭╮┃
┃┃┃┃┃┃╰╯┃╰┫╰┫┣━━┫╰━╯┃╰┫╰╯┣━━┃╰┫┃━┫┣━━┫┃┃╰┫╰╯┃╰╯┃┃━┫┃┃┃┃━┫╰╯┃
╰╯╰╯╰┻━━┻━┻━┻╯╱╱╰━━━┻━┻━━┻━━┻━┻━━┻╯╱╱╰╯╰━┻━━┻━━┻━━╯╰╯╰━━┻━━╯
`))
}
