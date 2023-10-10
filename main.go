package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Create configuration
	config := getConfig()

	// Create k8s client
	k8sclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting k8sclient, %s", err.Error())
	}

	// Create p4 client
	p4client, err := p4clientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting p4client, %s", err.Error())
	}

	// List nodes
	nodeList := getNodes(*k8sclient)
	for _, node := range nodeList.Items {
		fmt.Printf("%s\n", node.Name)
	}

	p4informers := p4informers.NewSharedInformerFactory(p4client, 10*time.Minute)

	c := newController(p4client, p4informers.P4kube().InternalVersion().P4s())

	channel := make(chan struct{})

	p4informers.Start(channel)
	c.run(channel)

}

func getConfig() *rest.Config {
	var kubeconfigpath *string

	// create filepath of kube config file which is at /home/apmec/.kube/config
	if home := homedir.HomeDir(); home != "" {
		kubeconfigpath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfigpath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	// creates configuration based on config path
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigpath)
	if err != nil {
		fmt.Printf("Could not get the config file due to %s", err.Error())

		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error %s, getting incluster config", err.Error())
		}

	}
	return config
}

func getNodes(k8sclient kubernetes.Clientset) *v1.NodeList {
	nodelist, _ := k8sclient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	return nodelist
}
