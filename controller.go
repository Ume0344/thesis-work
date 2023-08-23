package main

import (
	"fmt"
	p4clientset "p4kube/pkg/client/clientset/versioned"
	p4informers "p4kube/pkg/client/informers/internalversion/v1alpha1/internalversion"
	p4lister "p4kube/pkg/client/listers/v1alpha1/internalversion"
	"time"

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
	c.p4WorkQueue.Add(obj)
	fmt.Printf(" Deleting a P4 resource\n")
}

// Specifying the receiver of the method to be of type pointer to controller
// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) run(channel <-chan struct{}) {
	// Takes receive only channel

	// wait for the cache to be synched before starting workers
	if !cache.WaitForCacheSync(channel, c.p4Synched) {
		fmt.Print("Waiting for cache to be synched\n")
	}

	//Call the worker function after every 1 second till the channel is stopped
	fmt.Printf("Check\n")
	go wait.Until(c.worker, time.Second, channel)

	//Wait until some object is added into channel
	<-channel
}

func (c *Controller) worker() {

}
