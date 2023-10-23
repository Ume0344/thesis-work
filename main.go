package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	config := getConfig()

	k8sclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting k8sclient, %s", err.Error())
	}

	p4client, err := p4clientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting p4client, %s", err.Error())
	}

	p4s, err := p4client.P4kubeV1alpha1().P4s("p4-namespace").List(context.Background(), metav1.ListOptions{})

	for i, v := range p4s.Items {
		fmt.Printf("%d-%s\n", i, v.Name)
	}

	p4informers := p4informers.NewSharedInformerFactory(p4client, 10*time.Minute)

	c := newController(*k8sclient, p4client, p4informers.P4kube().InternalVersion().P4s())

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
