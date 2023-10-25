package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "p4kube/p4kube_grpc"
	"p4kube/pkg/apis/p4kube/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func checkStatus(finished chan bool, stream pb.P4DeploymentService_DeployP4Client) {
	for {
		status, _ := stream.Recv()
		log.Printf("Status: %v\n", status)
		if status != nil {
			if status.Status == "Deployed" || status.Status == "UnDeployed" {
				break
			}
		} else {
			time.Sleep(time.Second)
			continue
		}
	}
	finished <- true
}

func createCommand(p4resource v1alpha1.P4) string {
	cmdExec := fmt.Sprintf("cd %v; %v %v", p4resource.Spec.CompilerDirectory, p4resource.Spec.CompilerCommand, p4resource.Spec.P4Program)

	return cmdExec
}

func dialWorkerNode(p4resource v1alpha1.P4, address string) {
	// Connect to worker node
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Master node failed to connect to worker Node: %v", err.Error())
	}

	log.Println("Connection to worker node is successful")
	defer connection.Close()

	client := pb.NewP4DeploymentServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := client.DeployP4(ctx, &pb.P4Resource{Name: p4resource.Name, Command: createCommand(p4resource)})

	if err != nil {
		log.Fatalf("Could not receive response from server: %v", err)
	}

	finished := make(chan bool)

	go checkStatus(finished, stream)
	<-finished
}
