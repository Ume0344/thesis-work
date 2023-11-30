<center><h2><b>Week 14: 04.09 - 08.09.23</b></h2></center>

**P4 Custom Controllers**

- This week I did minor changes to controller;
- Taking compiler command from p4 manifest file.
- Measuring the time between when p4 resource is created and p4 program is deployed using t4p4s.
    - This time is on average 7 seconds.
- This is the [latest commit](https://git.comnets.net/p4-work/p4-kube/-/commit/ff24d251aa0a8ec6825baaaf3b13f5038e82d229) for the controller implementation so far. Please review it.

- I created a k8s cluster consists of 1 master node and 1 worker node. Ran the controller on the cluster. However, the controller could not deploy p4 resource on worker node because we donot have any scheduling mechanism for custom resources in k8s by default.

**Scheduling**

We have two options to schedule custom P4 resources on worker nodes.

**Implement a custom scheduler**

The steps involved could be;

*1- Node Discovery*: Discover all the available nodes where P4 is not being deployed.

*2- P4 Specification*: Define desired state of P4 in manifest files. The desired form could be successful deployment of P4 on node using t4p4s.

*3- Scheduling Algorithm / Scheduler*: Write scheduling algorithm to detect nodes where P4 is not deployed. Then schedule the P4 on any of the available nodes.

*4- Binding*: Binds the P4 to the node nominated by scheduling algorithm.

*5- Deploy the P4 on node*: Deploy P4 on the binded node through t4p4s.

**Pros**:
- We can customize the logic only for P4 resources.
- It can be a new contribution as I could not find any example where they are scheduling custom resources.

**Cons**

- High development work.
- Will require a lot of research and time.
- Does not have any implementations to get reference from.

**Deploy p4 controller on Pods and use default kube-scheduler for scheduling**

I need to research this option more in detail. So far, what I understood is following;

*1- Create another controller*: This controller will create a pod for our main controller who is handling a custom resource.

*2- Schedule the controller pod*: Schedule the controller pod on one of the available nodes through kube-scheduler.

**Pros**

- Comparitively lesser development work.
- We can refer kubevirt or knative how they deployed custom resources on pods.
- Scheduling can be done through k8s default scheduler.

**Cons**

- None, so far.
