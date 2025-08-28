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

package app

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/pflag"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"sigs.k8s.io/kubefed/pkg/controller/webhook/federatedtypeconfig"
	"sigs.k8s.io/kubefed/pkg/controller/webhook/kubefedcluster"
	"sigs.k8s.io/kubefed/pkg/controller/webhook/kubefedconfig"
	"sigs.k8s.io/kubefed/pkg/version"
)

var (
	certDir    string
	kubeconfig string
	masterURL  string
	port       = 8443
)

// NewWebhookCommand creates a *cobra.Command object with default parameters
func NewWebhookCommand(stopChan <-chan struct{}) *cobra.Command {
	verFlag := false

	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Start a kubefed webhook server",
		Long:  "Start a kubefed webhook server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stdout, "KubeFed webhook version: %s\n", fmt.Sprintf("%#v", version.Get()))
			if verFlag {
				os.Exit(0)
			}
			PrintFlags(cmd.Flags())

			if err := Run(stopChan); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flags.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flags.StringVar(&certDir, "cert-dir", "", "The directory where the TLS certs are located.")
	flags.IntVar(&port, "secure-port", port, "The port on which to serve HTTPS.")
	flags.BoolVar(&verFlag, "version", false, "Prints the Version info of webhook.")
	local := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	klog.InitFlags(local)
	flags.AddGoFlagSet(local)
	return cmd
}

// Run runs the webhook with options. This should never exit.
func Run(stopChan <-chan struct{}) error {
	logs.InitLogs()
	defer logs.FlushLogs()

	// Initialize the controller-runtime logger
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("error setting up webhook's config: %s", err)
	}
	webhookServer := webhook.NewServer(webhook.Options{
		// Port is the port number that the server will serve, Default: 9443
		Port: port,
		// Defaults to /tmp/k8s-webhook-server/serving-certs
		CertDir: certDir,
	})

	mgr, err := manager.New(config, manager.Options{
		WebhookServer: webhookServer,
	})
	if err != nil {
		klog.Fatalf("error setting up webhook manager: %s", err)
	}
	hookServer := mgr.GetWebhookServer()

	hookServer.Register("/validate-federatedtypeconfigs", &webhook.Admission{Handler: &federatedtypeconfig.AdmissionHook{}})
	hookServer.Register("/validate-kubefedcluster", &webhook.Admission{Handler: &kubefedcluster.AdmissionHook{}})
	hookServer.Register("/validate-kubefedconfig", &webhook.Admission{Handler: &kubefedconfig.Validator{}})
	hookServer.Register("/default-kubefedconfig", &webhook.Admission{Handler: &kubefedconfig.KubeFedConfigDefaulter{}})

	hookServer.WebhookMux().Handle("/readyz/", http.StripPrefix("/readyz/", &healthz.Handler{}))

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		klog.Fatalf("unable to run manager: %s", err)
	}

	return nil
}

// PrintFlags logs the flags in the flagset
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		klog.V(1).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
