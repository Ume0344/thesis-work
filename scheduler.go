package main

import (
	"context"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Scheduling UseCase1
func findRandomNode(k8sclient kubernetes.Clientset) v1.Node {
	nodeList, _ := k8sclient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	// Finding a random node to schedule p4 resource
	selectedNode := nodeList.Items[rand.Intn(len(nodeList.Items))]

	return selectedNode
}

func getNodeIpAddress(node v1.Node) string {
	ipAddress := node.Status.Addresses
	return ipAddress[0].Address
}
