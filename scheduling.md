**Scheduling in Kubernetes**

Kubernetes by default uses kube-scheduler for scheduling pods on nodes. It uses two approaches; filtering and scoring. In filtering, scheduler based on the pod requirements filters out the nodes which does not provide all the resources to the pods. 
After filtering out the nodes, it scores the available nodes based on scoring rules. And, the node that gets the higher score is scheduled for pod and kube-scheduler shares this information with API server.

**Scheduling Profiles**

A profile allows us to configure different stages of kube-scheduler.
By default (in default-scheduler profile), these are the [plugins](https://kubernetes.io/docs/reference/scheduling/config/#scheduling-plugins) which are enabled by kube-scheduler to schedule a pod.  To create a custom scheduler, we can enable or disable these plugins or even create our own plugins.

**[Kube-Virt](https://kubevirt.io/user-guide/operations/scheduler/) and Knative**

These use kube-scheduler with the default-scheduler profile. However, these allow users to use their own custom scheduler profiles. 

