package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion/v1alpha1/internalversion"
	p4lister "p4kube/pkg/client/listers/v1alpha1/internalversion"

	v1alpha1 "p4kube/pkg/apis/p4kube/v1alpha1"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	// Controller Struct which has attributes k8s standard clientset, p4 generated clientset
	// generated lister, cache and workqueue
	p4Client    p4clientset.Interface
	p4Lister    p4lister.P4Lister
	p4Synched   cache.InformerSynced //if cache has been synched with api server
	p4WorkQueue workqueue.RateLimitingInterface
}

func newController(
	p4Client p4clientset.Interface,
	p4Informer p4informers.P4Informer,

) *Controller {
	//Initialize the Controller struct and add event handler for registering
	//handler functions for adding and deleting p4 resources.

	c := &Controller{
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
	fmt.Printf("Creating a new P4 resource\n")
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
	fmt.Printf("Processing the items from queue\n")
	item, shutdown := c.p4WorkQueue.Get()

	// Delete the item from queue, so that we wont process it again
	defer c.p4WorkQueue.Forget(item)

	if shutdown {
		return false
	}

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
		fmt.Printf("Error getting p4resource %s\n", err.Error())
		return false
	}

	// %+v for printing struct
	fmt.Printf("p4 resource specs are :%+v\n", p4resource.Spec)

	c.handleP4Resource(p4resource.Spec)

	return true
}

// Handler Function to process created p4 spec
func (c *Controller) handleP4Resource(p4Spec v1alpha1.P4Spec) {
	//cmd := exec.Command("/bin/sh", "-c", "cd /home/$(whoami)/t4p4s; ./t4p4s.sh :l2fwd model=v1model")
	cmdExec := fmt.Sprintf("cd %v; ./t4p4s.sh %v model=v1model", p4Spec.CompilerDirectory, p4Spec.P4Program)
	cmd := exec.Command("/bin/sh", "-c", cmdExec)

	fmt.Print("Command to be executed : ", cmd, "\n")

	// Print output of command, also the error if command not successful.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("While running command, Getting error: %s\n", err.Error())
		fmt.Print("Cancelling the command\n")
		cmd.Cancel()
		// TODO: if command is not successfully executed, delete the created resource
		return
	}

	return
}
