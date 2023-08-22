package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	client "p4kube/pkg/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
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

	clientset, err := client.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting clientset, %s", err.Error())
	}

	p4, err := clientset.P4kubeV1alpha1().P4s("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("Cant get the list of P4s due to %s", err.Error())
	}

	for i, v := range p4.Items {
		fmt.Printf("p4-%d: %v \n", i, v.Name)
	}
}
