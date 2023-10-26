package main

import (
	"context"
	"log"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Scheduling UseCase1
func findRandomNode(k8sclient kubernetes.Clientset) v1.Node {
	var workerNodeList []v1.Node
	var masterFlag bool

	nodeList, _ := k8sclient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	for _, node := range nodeList.Items {
		for label := range node.Labels {
			masterFlag = false
			if label == "node-role.kubernetes.io/control-plane" {
				masterFlag = true
				break
			}
		}
		if !masterFlag {
			workerNodeList = append(workerNodeList, node)
		}
	}

	log.Println("Printing avaialable worker nodes")
	for i, node := range workerNodeList {
		log.Printf("worker node %d: %s", i+1, node.Name)
	}

	// Finding a random node to schedule p4 resource
	selectedNode := workerNodeList[rand.Intn(len(workerNodeList))]

	return selectedNode
}

func getNodeIpAddress(node v1.Node) string {
	ipAddress := node.Status.Addresses
	return ipAddress[0].Address
}
