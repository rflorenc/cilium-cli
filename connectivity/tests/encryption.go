// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package tests

import (
	"context"
	"fmt"

	"github.com/cilium/cilium-cli/connectivity/check"
)

// OutsideToNodePort sends an HTTP request from client pod running on a node w/o
// Cilium to NodePort services.
func EncryptPodToPod() check.Scenario {
	return &encryptPodToPod{}
}

type encryptPodToPod struct{}

func (s *encryptPodToPod) Name() string {
	return "encrypt-pod-to-pod"
}

func (s *encryptPodToPod) Run(ctx context.Context, t *check.Test) {
	// pod1 @ node1
	// pod2 @ node2
	// derive iface
	// tcpdump @ host netns node1
	// tcpdump @ host netns node2
	// run curl

	clientPod := t.Context().HostNetNSPodsByNode()[t.NodesWithoutCilium()[0]]
	i := 0

	for _, svc := range t.Context().EchoServices() {
		for _, node := range t.Context().CiliumPods() {
			node := node // copy to avoid memory aliasing when using reference

			curlNodePort(ctx, s, t, fmt.Sprintf("curl-%d", i), &clientPod, svc, &node)
			i++
		}
	}

}
