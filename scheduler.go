package main

import (
	"context"
	"log"
	"math/rand"
	"p4kube/pkg/apis/p4kube/v1alpha1"
	"p4kube/pkg/client/clientset/versioned"

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

func scheduleSplittedResource(k8sclient kubernetes.Clientset, p4resource *v1alpha1.P4, p4client versioned.Interface) (v1.Node, string) {
	var node v1.Node
	var nodeName string
	var command string = p4resource.Spec.CompilerCommand
	var flag bool

	p4Conversion := "p4-conversion"
	cCompilation := "c-compilation"
	switchCompilation := "switch-compilation"

	// Check if user mentioned deployment phase in manifest file. If deployment phase is p4-conversion,
	// find a random node and deploy all phases onto it. If deployment phase is c-compilation, find
	// the node where p4-conversion is already executed. If scheduler cant find any node where p4-conversion is present,
	// find random node and just deploy the c compilation. If deployment phase is switch-compilation,
	// find the node where c-compilation is already executed for that particular
	// network function. If scheduler cant find any node where c-compilation is present, check if p4conversion
	// is present, if not find random node and just deploy the switch compilation.
	if p4resource.Spec.DeploymentPhase == p4Conversion {
		log.Printf("Deployment Phase is P4 to C conversion, finding a random node")
		node = findRandomNode(k8sclient)

	} else if p4resource.Spec.DeploymentPhase == cCompilation {
		log.Printf("Deployment Phase is C compilation, finding a node where build files are already present")
		flag, nodeName = checkP4ConversionExists(p4resource, p4client, p4Conversion)
		if flag {
			log.Printf("P4 conversion for resource %v already executed", p4resource.Name)
			node = findTargetNode(k8sclient, nodeName)
		} else {
			log.Printf("P4 conversion does not exist on any node, finding a random node to deploy c compilation")
			node = findRandomNode(k8sclient)
			command = "./t4p4s.sh p4 c model=v1model"
		}
	} else if p4resource.Spec.DeploymentPhase == switchCompilation {
		log.Printf("Deployment Phase is switch compilation, finding a node where build files are already present")
		flag, nodeName = checkCompilationExists(p4resource, p4client, cCompilation)
		if flag {
			log.Printf("C compilation for resource %v already executed", p4resource.Name)
			node = findTargetNode(k8sclient, nodeName)
		} else {
			log.Printf("C compilation does not exist on any node, checking if p4 conversion exists")
			flag, nodeName = checkP4ConversionExists(p4resource, p4client, p4Conversion)
			if flag {
				log.Printf("P4 conversion for resource %v already executed", p4resource.Name)
				node = findTargetNode(k8sclient, nodeName)
				command = "./t4p4s.sh c run model=v1model"
			} else {
				node = findRandomNode(k8sclient)
				command = "./t4p4s.sh p4 c run model=v1model"
			}
		}
	}
	return node, command
}

// Check if P42C conversion for the required p4resource already executed, if yes,
// return the node name where P42C conversion happened
func checkP4ConversionExists(p4resource *v1alpha1.P4, p4client versioned.Interface, deploymentPhase string) (bool, string) {
	flag := false
	var node string

	p4s, _ := p4client.P4kubeV1alpha1().P4s("p4-namespace").List(context.Background(), metav1.ListOptions{})
	for _, p4 := range p4s.Items {
		if p4.Status.NetworkFunction == p4resource.Spec.NetworkFunction && p4.Status.DeploymentPhase == deploymentPhase {
			node = p4.Status.Node
			log.Printf("Found the node where build files for p4 resource '%v' are present: %v", p4resource.Name, node)
			flag = true
			break
		}
	}

	return flag, node
}

// Check if C compilation for the required p4resource already executed, if yes,
// return the node name where c compilation happened
func checkCompilationExists(p4resource *v1alpha1.P4, p4client versioned.Interface, deploymentPhase string) (bool, string) {
	flag := false
	var node string

	p4s, _ := p4client.P4kubeV1alpha1().P4s("p4-namespace").List(context.Background(), metav1.ListOptions{})
	for _, p4 := range p4s.Items {
		if p4.Status.NetworkFunction == p4resource.Spec.NetworkFunction && p4.Status.DeploymentPhase == deploymentPhase {
			node = p4.Status.Node
			log.Printf("Found the node where build files for p4 resource '%v' are present: %v", p4resource.Name, node)
			flag = true
			break
		}
	}

	return flag, node
}
