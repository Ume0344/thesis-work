package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func findNodes(k8sclient kubernetes.Clientset) {
	nodeList, _ := k8sclient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	for _, node := range nodeList.Items {
		fmt.Printf("%s\n", node.Name)
	}
}
