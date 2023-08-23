package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion"

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

	p4client, err := p4clientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error getting p4client, %s", err.Error())
	}

	p4informers := p4informers.NewSharedInformerFactory(p4client, 10*time.Minute)

	c := newController(p4client, p4informers.P4kube().InternalVersion().P4s())

	channel := make(chan struct{})

	p4informers.Start(channel)
	c.run(channel)
}
