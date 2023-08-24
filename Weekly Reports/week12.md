<center><h2><b>Week 12: 21.08 - 25.08.23</b></h2></center>

**P4 Custom Controllers**

- This week I started implemeneting P4 custom controller. Till now, controller can;
    - Implements the code-generator to create clientset, informers and listers for custom resources.
    - Access the P4 custom resources.
    - Start the informer to keep track of whenever a p4 resource is created or deleted.
    - Get the objects from worker queue.
    - This is the [latest commit](https://git.comnets.net/p4-work/p4-kube/-/commit/52e609b21b50060afec7c897c05ceebc1cb969c0) for the controller implementation so far. Please review it.

- Next week, I will keep working on custom controller to handle the created or deleted p4 resources.