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
	"os"
	"runtime"
	"strconv"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	_ "go.uber.org/automaxprocs"
	"sigs.k8s.io/kubefed/pkg/constants"
)

// Base version information.
//
// This is the fallback data used when version information from git is not
// provided via go ldflags (via Makefile).
var (
	Version = "latest" // output of "git describe"
	// GitCommit the prerequisite is that the branch should be
	// tagged using the correct versioning strategy.

	GitCommit    = "unknown" // sha1 from git, output of $(git rev-parse HEAD)
	GitTreeState = "unknown" // state of a git tree, either "clean" or "dirty"

	BuildDate = "unknown" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')

	Community = constants.DefaultCommunity
	Author    = constants.DefaultAuthor
)

type Info struct {
	Community    string `json:"community"`
	Author       string `json:"author"`
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
		Community:    Community,
		Author:       Author,
		Version:      Version,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
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

// Print the version information.
func Print() {
	v := Get()
	// 创建表格
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// 设置表头（横向）
	t.AppendHeader(table.Row{
		"Community", "Author", "Version", "Git Commit", "Build Date",
		"Go Version", "Compiler", "Platform", "Runtime Cores", "Total Memory",
	})

	// 添加数据
	t.AppendRow([]interface{}{
		v.Community, v.Author, v.Version, v.GitCommit,
		v.BuildDate, v.GoVersion, v.Compiler, v.Platform,
		strconv.Itoa(v.RuntimeCores) + " cores",
		strconv.Itoa(v.TotalMem) + " KB",
	})

	// 设置表格样式
	t.SetStyle(table.StyleDefault)                      // 轻量风格
	t.Style().Format.Header = text.FormatUpper          // 加粗表头
	t.Style().Color.Header = text.Colors{text.FgHiBlue} // 高亮蓝色
	t.Style().Options.SeparateRows = true               // 让每一行独立，更清晰

	// 渲染表格
	t.Render()
}

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
