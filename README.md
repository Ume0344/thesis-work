**Create ClientSet, Informers, Listers through [Code Generator](https://github.com/kubernetes/code-generator.git)**

1- Create types.go
- This file will have the definition of P4 object and P4 Specs.

2- Create register.go
- This file will register the P4 type to kubernetes scheme.

3- We also need doc.go file where we define the global tags (tags are the way to control the behavior of code-generators)
- Tags could be: Local and Global tags

4- These three files are needed to run code-generator which then generates the following resources;
- ClientSet - To interact with custom resources.

- Lister - To get the objects from informers.

- Informers - To keep track of created and deleted resources in API servers.

- DeepCopy Objects - To deepcopy objects, to register struct to kubernetes as k8s objects

5- Get the code-generators in home/apmec/go/k8s.io by; 

`git clone https://github.com/kubernetes/code-generator.git`

6- Set; 

`execDir=/home/apmec/go/src/k8s.io/code-generator/`

7- Run this command from project root directory; 

`$execDir/generate-groups.sh deepcopy,client,informer,lister p4kube/pkg/client p4kube/pkg/apis p4kube:v1alpha1 -h $execDir/examples/hack/boilerplate.go.txt --output-base ..`

It will generate all the files (clientset, informers, lister etc in ./pkg ).
