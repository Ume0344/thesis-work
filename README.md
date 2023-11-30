**Improving Performance of Orchestration for Cloud-Native Applications Using Programmable Network Devices**

**Abstract**

In  recent  years,  the  orchestration  of  cloud-native  applications  has  experienced  tremendous 
popularity. Kubernetes, an open-source orchestration engine, orchestrates containerized 
applications by managing deployment, scalability, and failure recovery. As Kubernetes thrives in 
managing containerized workloads, several frameworks have emerged to extend its capabilities. 
For instance, KubeVirt extends Kubernetes to manage not only containerized workloads but also 
virtual machines (VMs). Similarly, Knative facilitates the orchestration of serverless workloads. 
However, efficient communication between containerized applications within a Kubernetes cluster 
comes with a challenge. In containerized network functions, network packets as they traverse the 
network stack experience increased latency, moving from user-space to kernel-space before being 
deployed onto the network interface card (NIC). This added latency can affect the responsiveness 
and performance of applications, particularly those with strict latency requirements. Additionally, 
containerized  network  functions  deployed  within  a  Kubernetes  cluster  are  typically  not  easily 
programmable  at  runtime.  This  inflexibility  can  be  limiting  in  scenarios  requiring  dynamic 
network  configuration  and  adaptation.  To  address  these  challenges,  we  introduce  P4Kube,  a 
framework that bridges the gap between orchestrating containerized workloads in Kubernetes and 
the dynamic programmability of network functions. P4Kube leverages P4 (Protocol-independent 
packet processor), a domain-specific programming language designed for runtime programming 
of network devices. This framework will enable the deployment of network functions written in 
P4 language with the containerized applications on a multi-node Kubernetes cluster. Furthermore, 
P4Kube  introduces  an  intelligent  scheduler  to  schedule  the  network  functions  on  multi-node 
cluster based on node deployment conditions. It also makes use of a P4 compiler that works with 
emerging  hardware  acceleration  technologies  like  the  Data  Plane  Development  Kit  (DPDK)  to 
enhance packet-processing efficiency. By reducing packet-processing time and provisioning time, 
P4Kube aims to provide network communication with minimal latency, creating a more responsive 
and efficient environment for containerized applications within a Kubernetes cluster.
