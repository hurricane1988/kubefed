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

package leaderelection

import (
	"context"
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"

	"sigs.k8s.io/kubefed/cmd/controller-manager/app/options"
)

const (
	component = "kubefed-controller-manager"
	userAgent = "kubefed-leader-election"
)

func NewKubeFedLeaderElector(opts *options.Options, fnStartControllers func(*options.Options, <-chan struct{})) (*leaderelection.LeaderElector, error) {
	kubeConfig := restclient.CopyConfig(opts.Config.KubeConfig)
	// 在 rest.Config 中添加一个 User-Agent 为kubefed-leader-election标头,用于标识客户端
	restclient.AddUserAgent(kubeConfig, userAgent)
	// 初始化一个kubernetes的Clientset
	leaderElectionClient := kubeclient.NewForConfigOrDie(kubeConfig)

	hostname, err := os.Hostname()
	if err != nil {
		klog.Infof("unable to get hostname: %v", err)
		return nil, err
	}

	// 用于创建一个事件广播器 EventBroadcaster 实例，它可以用来将事件分发给多个事件接收者
	// 1. 记录事件：将事件写入 Kubernetes API 服务器，以便通过 kubectl describe 等工具查看
	// 2. 输出到日志：将事件打印到日志文件或标准输出中，便于调试和监控
	broadcaster := record.NewBroadcaster()

	// StartRecordingToSink 该方法将事件广播到指定的接收器（Sink）
	// 参数是 EventSinkImpl 的实例，EventSinkImpl 是 EventSink 接口的标准实现，用于将事件发送到 Kubernetes API 服务器
	// leaderElectionClient.CoreV1().Events(...) 用来访问 Kubernetes 的事件 API
	// 指定在特定的命名空间中记录事件（opts.Config.KubeFedNamespace）
	broadcaster.StartRecordingToSink(&corev1.EventSinkImpl{Interface: leaderElectionClient.CoreV1().Events(opts.Config.KubeFedNamespace)})

	// 生成一个组件kubefed-controller-manager的事件记录
	eventRecorder := broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: component})

	// 添加一个唯一标识符，以确保同一主机上的两个进程不会意外地都变为活跃状态
	id := hostname + "_" + string(uuid.NewUUID())

	// resourcelock.New 用于生成一个新的资源锁，确保在多个实例中只有一个被选为领导者
	rl, err := resourcelock.New(
		string(opts.LeaderElection.ResourceLock), // 锁类型（分为：configmaps、endpoints、leases 三种）
		opts.Config.KubeFedNamespace,             // 锁所在的命名空间
		component,                                // 组件名称，此处为kubefed-controller-manager
		leaderElectionClient.CoreV1(),            // Kubernetes client 用于操作 CoreV1（通常是 Pod、Service 等资源）
		leaderElectionClient.CoordinationV1(),    // Kubernetes client 用于操作 CoordinationV1（用于领导者选举）

		// 锁的配置
		resourcelock.ResourceLockConfig{
			Identity:      id,            // 当前实例的唯一身份标识，此处为<hostname>_<
			EventRecorder: eventRecorder, // 事件记录器，用于记录锁的变更或事件
		})
	if err != nil {
		klog.Infof("couldn't create resource lock: %v", err)
		return nil, err
	}

	// 根据LeaderElectionConfig自定义配置生成一个LeaderElector对象
	return leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		// 资源锁对象,确保多个实例竞争同一资源时，同一时刻只有一个实例能持有锁
		Lock: rl,
		// 租约持续时间（time.Duration），定义领导者保持锁的最长时间
		LeaseDuration: opts.LeaderElection.LeaseDuration,
		// 续约期限（time.Duration），定义领导者在锁定资源后必须在此时间内成功续约，否则视为失去领导权
		RenewDeadline: opts.LeaderElection.RenewDeadline,
		// 续约期限（time.Duration），定义领导者在锁定资源后必须在此时间内成功续约，否则视为失去领导权
		RetryPeriod: opts.LeaderElection.RetryPeriod,

		// 领导者选举的回调函数，定义在不同选举事件中触发的逻辑
		Callbacks: leaderelection.LeaderCallbacks{

			// 在当前实例成为领导者时触发
			OnStartedLeading: func(ctx context.Context) {
				klog.Info("promoted as leader")
				// 等待 ctx.Done() 通知来清理和退出
				stopChan := ctx.Done()
				// 调用 fnStartControllers 启动控制器逻辑
				fnStartControllers(opts, stopChan)
				<-stopChan
			},
			// 当当前实例失去领导权时触发
			OnStoppedLeading: func() {
				klog.Info("leader election lost")
			},
		},
	})
}
