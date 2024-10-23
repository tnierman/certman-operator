/*
main defines the local
*/

package main

import (
	"fmt"
	"os"

	minikube "k8s.io/minikube/cmd/minikube/cmd"
)

const (
	kubeVersion = "v1.21.5"
)

var (
	args = []string{"start", "--memory=6g", "--bootstrapper=kubeadm", "--extra-config=kubelet.authentication-token-webhook=true", "--extra-config=kubelet.authorization-mode=Webhook", "--extra-config=scheduler.address=0.0.0.0", "--extra-config=controller-manager.address=0.0.0.0", "--extra-config=apiserver.service-node-port-range=1-65535"}
)

func main() {
	cmd := minikube.RootCmd
	args = append(args, kubeVersion)
	//args = []string{"podman-env", "--help"}
	cmd.SetArgs(args)

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute minikube: %v", err)
		os.Exit(1)
	}
}

