**Scheduling in Kubernetes**

Kubernetes by default uses kube-scheduler for scheduling pods on nodes. It uses two approaches; filtering and scoring. In filtering, scheduler based on the pod requirements filters out the nodes which does not provide all the resources to the pods. 

e.g, `VolumeBinding` is an example of filter predicate which filters out the nodes which cannot bind the requested volume.

After filtering out the nodes, it scores the available nodes based on scoring rules. And, the node that gets the higher score is scheduled for pod and kube-scheduler shares this information with API server.

e.g, `ImageLocality` is a scoring predicate. The nodes get the higher score which already has the container image locally.

`NodeResourcesBalancedAllocation` is another scoring predicate which gives higher priority to the nodes which are underutilized.

**Scheduling Profiles**

A profile allows us to configure different stages of kube-scheduler.
By default (in default-scheduler profile), these are the [plugins](https://kubernetes.io/docs/reference/scheduling/config/#scheduling-plugins) which are enabled by kube-scheduler to schedule a pod.  To create a custom scheduler, we can enable or disable these plugins or even create our own plugins.

**[Kube-Virt](https://kubevirt.io/user-guide/operations/scheduler/) and Knative**

These use kube-scheduler with the default-scheduler profile. However, these allow users to use their own custom scheduler profiles. 

KuberVirt supports the scheduling of VMs based on k8s nodeSelector or nodeAffinity rules. 

**nodeSelector**

KubeVirt can schedule a VM on specific node based on `nodeSelector` label defined in manifest file. Forexample, we have a node with label; `performance: high`. We can use `nodeSelector` label for a VM to be deployed on only those nodes which have the label `performance: high`.

**nodeAffinity Rules**

There are two types of affinity that can be applied to nodes and pods.

- **requiredDuringSchedulingIgnoredDuringExecution**

- This is the hard requirement affinity. Scheduler only schedules the VM on node which fullfills this requirement.

- e.g, Scheduler *only* schedules the VM on node with label `performance: high`
```
affinity:
        nodeAffinity: 
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: performance
                    operator: In
                    values:
                      - high
```

- **preferredDuringSchedulingIgnoredDuringExecution**

- This is the soft requirement affinity. Schedular prefers to schedule the VM based on this rule but if it does not find node with specific requirement, it will schedule it to other node.

- e.g, 
The specific example prefers to schedule VM on node having label; `performance: high` but if it does not find any node having this label, it will schedule it to other node.

```
affinity:
        nodeAffinity: 
          preferredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: performance
                    operator: In
                    values:
                      - high
```

**References**

https://kubevirt.io/2020/Advanced-scheduling-with-affinity-rules.html
https://community.ops.io/danielepolencic/kubernetes-scheduler-deep-dive-khc


