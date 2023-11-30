<center><h2><b>Week 21: 23.10 - 27.10.23</b></h2></center>

- This week, I implemented grpc server and client to deploy p4 resource from master node to worker node.
- Setup the k8s cluster of one master node and 3 worker nodes.
- Installed t4p4s on them.
- Scheduling Usecase 1, which is to randomly select a node and deploy p4 resource, is completed.


**GRPC Communication**

- Install protobuf compiler by running these commands;
```
$PB_REL="https://github.com/protocolbuffers/protobuf/releases"
$ curl -LO $PB_REL/download/v3.12.1/protoc-3.12.1-linux-x86_64.zip

$sudo apt install unzip
$unzip protoc-3.12.1-linux-x86_64.zip -d HOME/.local

$export PATH="$PATH:$HOME/.local/bin"
```

- Install protocol compiler plugins
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- To regenerate gRPC code, Run;
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```