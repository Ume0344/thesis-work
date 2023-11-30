- Download go; 
`wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz`
- Remove any previous go installations;
`sudo apt-get remove golang-go && sudo rm -rf /usr/local/go/`
- Install go; 
`sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz &&export PATH=$PATH:/usr/local/go/bin`
- Go to project directory and get the dependencies;
`cd /home/$(whoami)/p4kube && go get .` 
- Create a namespace for p4 resources;
`kubectl create namespace p4-namespace`
- Create CRD for p4 resource;
`kubectl create -f home/$(whoami)/p4kube/manifests/p4-crds.yaml`
- Create a p4 resource;
`Kubectl create -f home/$(whoami)/p4kube/manifests/p4-cr.yaml` 
- To list created p4resources;
 `kubectl get p4s -n p4-namespace`
- To see details of a p4 resource;
 `kubectl get p4s -n p4-namespace <p4_resource_name> -o yaml`
- To create a p4 manifest file, please refer to examples in manifests folder.
