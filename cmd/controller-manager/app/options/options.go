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

// Package options contains flags and options for initializing controller-manager
package options

import (
	"time"

	"github.com/spf13/pflag"
	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"sigs.k8s.io/kubefed/pkg/controller/utils"
)

// Options contains everything necessary to create and run controller-manager.
type Options struct {
	Config                   *utils.ControllerConfig
	FeatureGates             map[string]bool
	Scope                    apiextv1.ResourceScope
	LeaderElection           *utils.LeaderElectionConfiguration
	ClusterHealthCheckConfig *utils.ClusterHealthCheckConfig
}

// AddFlags adds flags to fs and binds them to options.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	// KubeFed 基础配置
	fs.StringVar(&o.Config.KubeFedNamespace, "kubefed-namespace", "", "The namespace the KubeFed control plane is deployed in.")
	// Leader 选举参数绑定
	fs.DurationVar(&o.LeaderElection.LeaseDuration, "leader-elect-lease-duration", 15*time.Second,
		"The maximum duration that a leader can be stopped before it is replaced by another candidate.")
	fs.DurationVar(&o.LeaderElection.RenewDeadline, "leader-elect-renew-deadline", 10*time.Second,
		"The interval between attempts by the acting master to renew a leadership slot before it stops leading.")
	fs.DurationVar(&o.LeaderElection.RetryPeriod, "leader-elect-retry-period", 5*time.Second,
		"The duration the clients should wait between attempting acquisition and renewal of a leadership.")
	// 绑定资源锁类型 (例如: "leases", "configmaps", "endpoints")
	fs.StringVar((*string)(&o.LeaderElection.ResourceLock), "leader-elect-resource-lock", fedv1b1.LeasesResourceLock,
		"The type of resource object that will be used to lock during leader election.")
	// 其他配置
	if o.ClusterHealthCheckConfig != nil {
		fs.DurationVar(&o.ClusterHealthCheckConfig.Period, "cluster-health-check-period", 10*time.Second, "How often to check the health of cluster.")
	}
}

func NewOptions() *Options {
	return &Options{
		Config:                   new(utils.ControllerConfig),
		FeatureGates:             make(map[string]bool),
		LeaderElection:           new(utils.LeaderElectionConfiguration),
		ClusterHealthCheckConfig: new(utils.ClusterHealthCheckConfig),
	}
}
