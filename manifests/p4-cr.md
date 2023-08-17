**Guide to create P4 Custom Resource**

1- Create a .yaml file and use p4-cr.yaml file as reference.

2- Please add `spec.p4ProgramLocation` and `spec.compilerCommand`, these are the required arguments. Otherwise custom resource will not be created.

3- Run the following command to create p4 custom resource;

```
kubectl apply -f <cr file name, i.e, p4-cr.yaml>
```
