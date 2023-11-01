package main

import (
	"context"
	"log"
	"math/rand"
	"p4kube/pkg/apis/p4kube/v1alpha1"
	"p4kube/pkg/client/clientset/versioned"

	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Scheduling UseCase1
func findRandomNode(k8sclient kubernetes.Clientset) v1.Node {
	workerNodeList := getWorkerNodeList(k8sclient)
	// Finding a random node to schedule p4 resource
	selectedNode := workerNodeList[rand.Intn(len(workerNodeList))]

	return selectedNode
}

func findTargetNode(k8sclient kubernetes.Clientset, targetNode string) v1.Node {
	workerNodeList := getWorkerNodeList(k8sclient)
	var targetNodeSpec v1.Node

	for _, node := range workerNodeList {
		if node.Name == targetNode {
			targetNodeSpec = node
			break
		}
	}

	return targetNodeSpec
}

func getNodeIpAddress(node v1.Node) string {
	ipAddress := node.Status.Addresses
	return ipAddress[0].Address
}

func getWorkerNodeList(k8sclient kubernetes.Clientset) []v1.Node {
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

	log.Println("Avaialable worker nodes")
	for i, node := range workerNodeList {
		log.Printf("worker node %d: %s", i+1, node.Name)
	}

	return workerNodeList
}

func scheduleSplittedResource(k8sclient kubernetes.Clientset, p4resource *v1alpha1.P4, p4client versioned.Interface) v1.Node {
	var node v1.Node
	var nodeName string
	var flag bool

	p42c := " p4 "
	cCompilation := " c "

	// Check if compiler command has 'p4', if yes,. find a random node for it to be deployed.
	// else if compiler command has 'c' in it, check if p4 conversion already executed. If conversion is already executed,
	// get the node where conversion happend and schedule c compilation on that node
	if strings.Contains(p4resource.Spec.CompilerCommand, p42c) {
		log.Printf("Deployment Phase is P4 to C conversion, finding a random node")
		node = findRandomNode(k8sclient)

	} else if strings.Contains(p4resource.Spec.CompilerCommand, cCompilation) {
		log.Printf("Deployment Phase is C compilation, finding a node where build files are already present")
		flag, nodeName = checkP4ConversionExists(p4resource, p4client)
		if flag {
			log.Printf("P4 conversion for resource %v already executed", p4resource.Name)
			node = findTargetNode(k8sclient, nodeName)
		} else {
			log.Fatalf("Could not found P4 conversion for resource %v. Please first run p4 to c conversion", p4resource.Name)
		}
	}

	return node
}

// Check if P42C conversion for the required p4resource already executed, if yes,
// return the node name where P42C conversion happened
func checkP4ConversionExists(p4resource *v1alpha1.P4, p4client versioned.Interface) (bool, string) {
	flag := false
	var node string

	newP4Name := strings.Split(p4resource.Name, "-")[0]

	p4s, _ := p4client.P4kubeV1alpha1().P4s("p4-namespace").List(context.Background(), metav1.ListOptions{})
	for _, p4 := range p4s.Items {
		if strings.Contains(p4.Name, "p4") && strings.Contains(p4.Name, newP4Name) {
			node = p4.Status.Node
			log.Printf("Found the node where build files for p4 resource '%v' are present: %v", newP4Name, node)
			flag = true
			break
		}
	}

	return flag, node
}
