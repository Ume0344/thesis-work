package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion/v1alpha1/internalversion"
	p4lister "p4kube/pkg/client/listers/v1alpha1/internalversion"

	v1alpha1 "p4kube/pkg/apis/p4kube/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	port = "50051"
)

type Controller struct {
	// Controller Struct which has attributes k8s standard clientset, p4 generated clientset
	// generated lister, cache and workqueue
	k8sclient   kubernetes.Clientset
	p4Client    p4clientset.Interface
	p4Lister    p4lister.P4Lister
	p4Synched   cache.InformerSynced //if cache has been synched with api server
	p4WorkQueue workqueue.RateLimitingInterface
}

func newController(
	k8sclient kubernetes.Clientset,
	p4Client p4clientset.Interface,
	p4Informer p4informers.P4Informer,

) *Controller {
	//Initialize the Controller struct and add event handler for registering
	//handler functions for adding and deleting p4 resources.

	c := &Controller{
		k8sclient:   k8sclient,
		p4Client:    p4Client,
		p4Lister:    p4Informer.Lister(),
		p4Synched:   p4Informer.Informer().HasSynced,
		p4WorkQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "p4kube"),
	}

	p4Informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.handleAdd,
		DeleteFunc: c.handleDel,
	})

	return c
}

func (c *Controller) handleAdd(obj interface{}) {
	// Add objects to queue
	c.p4WorkQueue.Add(obj)
	fmt.Printf("Handling a P4 resource\n")
}

func (c *Controller) handleDel(obj interface{}) {
	c.p4WorkQueue.Done(obj)
	fmt.Printf("Deleting a P4 resource\n")
}

// Specifying the receiver of the method to be of type pointer to controller
// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) run(channel <-chan struct{}) {
	// Takes receive only channel

	// wait for the cache inside the informer to be synched before starting workers
	if !cache.WaitForCacheSync(channel, c.p4Synched) {
		fmt.Print("Waiting for cache to be synched\n")
	}

	//Create goroutine to call the worker function after every 1 second till the channel is stopped
	go wait.Until(c.worker, time.Second, channel)

	//Wait until some object is added into channel
	<-channel
}

func (c *Controller) worker() {
	// loop till processItem returns true, on false it will wait for a second and then again this function will be called by run()
	for c.processNextItem() {

	}
}

func (c *Controller) processNextItem() bool {
	// process the items from queue
	fmt.Printf("Processing the items from queue %v\n", c.p4WorkQueue.Len())
	item, shutdown := c.p4WorkQueue.Get()

	// Delete the item from queue, so that we wont process it again
	defer c.p4WorkQueue.Forget(item)

	if shutdown {
		return false
	}

	startTime := time.Now()
	// Generating key for each item in queue
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Printf("Getting key from cache %s\n", err.Error())
		return false
	}

	// Getting namespace, name from genrated key
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Printf("Getting namespace and name from MetaNamespaceKeyFunc %s\n", err.Error())
		return false
	}

	p4resource, err := c.p4Lister.P4s(ns).Get(name)

	if err != nil {
		fmt.Printf("Error getting P4 resource %s\n", err.Error())
		return false
	}

	// %+v for printing struct
	fmt.Printf("P4 resource specs are :%+v\n", p4resource.Spec)

	// Get the item, check its status, if status is deployed, forget it and
	// donot call handleP4Resource
	deployment := c.handleP4Resource(p4resource, startTime)

	if deployment {
		c.p4WorkQueue.Forget(item)
	}

	return true
}

// Handler Function to process created p4 spec
func (c *Controller) handleP4Resource(p4resource *v1alpha1.P4, startTime time.Time) bool {
	var deploy bool
	var selectedNode v1.Node

	if p4resource.Status.Progress == "Deployed" {
		fmt.Printf("P4 resource %s already deployed. Removing it from the queue.\n", p4resource.Name)
		deploy = true
		return deploy

	} else {
		// 'split' is a check if user wants p4 deployment in phases
		split := c.splitCheck(p4resource.Spec.CompilerCommand)
		if split {
			log.Println("Deployment is in 3 phases")
			selectedNode = scheduleSplittedResource(c.k8sclient, p4resource, c.p4Client)
		} else if p4resource.Spec.TargetNode == "" {
			log.Println("Target Node is not mentioned in p4 manifest file, scheduling it to a random available node")
			selectedNode = findRandomNode(c.k8sclient)
		} else if p4resource.Spec.TargetNode != "" {
			log.Printf("Target Node mentioned in p4 manifest file: %s", p4resource.Spec.TargetNode)
			selectedNode = findTargetNode(c.k8sclient, p4resource.Spec.TargetNode)
		}

		log.Printf("Scheduling P4 resource to Node: %v", selectedNode.Name)

		nodeIpAddress := getNodeIpAddress(selectedNode)
		log.Printf("Connecting to worker node at IP : %s", nodeIpAddress)

		nodeAddressWithPort := fmt.Sprintf("%s:%s", nodeIpAddress, port)
		log.Printf("Address : %v", nodeAddressWithPort)

		deploymentStatus := dialWorkerNode(p4resource, nodeAddressWithPort)
		deploy = c.updateP4Status(p4resource, deploymentStatus, selectedNode)

		stopTime := time.Since(startTime)

		fmt.Printf("Provisioning time for P4 Resource: %v\n", stopTime)
	}

	return deploy
}

// Updates the P4 status after an attempt for its deployment
func (c *Controller) updateP4Status(p4resource *v1alpha1.P4, deploymentStatus string, node v1.Node) bool {
	var deploy bool
	if deploymentStatus == "Deployed" {
		p4resource.Status.Progress = "Deployed"
		p4resource.Status.Node = node.Name
		p4resource.Status.NetworkFunction = p4resource.Spec.NetworkFunction
		p4resource.Status.DeploymentPhase = p4resource.Spec.DeploymentPhase
		deploy = true
	} else {
		p4resource.Status.Progress = "Deployment Unsuccessful"
		fmt.Printf("P4 resource %s could not be deployed, deleting the resource\n", p4resource.Name)
		fmt.Print("Cancelling the command\n")
		// cmd.Cancel()
		// TODO: if command is not successfully executed, delete the created resource
		deploy = false
	}

	_, err := c.p4Client.P4kubeV1alpha1().P4s(p4resource.Namespace).UpdateStatus(context.Background(), p4resource, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("Error updating the p4 resource: %s", err.Error())
	}

	return deploy
}

// Check if user wants deployment in phases based on compiler command mentioned in manifest file
func (c *Controller) splitCheck(compilerCommand string) bool {
	var flag bool

	if strings.Contains(compilerCommand, " p4 ") || strings.Contains(compilerCommand, " c ") || strings.Contains(compilerCommand, " run ") {
		flag = true
	} else {
		flag = false
	}

	return flag
}
