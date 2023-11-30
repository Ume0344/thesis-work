<center><h2><b>Week 13: 28.09 - 01.09.23</b></h2></center>

**P4 Custom Controllers**

- This week I continued implemeneting P4 custom controller. Now, controller can;
    - Implements the code-generator to create clientset, informers and listers for custom resources.
    - Accesses the P4 custom resources.
    - Starts the informer to keep track of whenever a p4 resource is created or deleted.
    - Gets the objects from worker queue.
    - Handles the created P4 resources.
    - Updates the status of P4 resource, if deployment gets successful.
    - This is the [latest commit](https://git.comnets.net/p4-work/p4-kube/-/commit/da4dbd42e7c76af535b21de056ea67aae6b630b5) for the controller implementation so far. Please review it.

- Next week, we will test this controller on our cluster instead of minikube cluster.
