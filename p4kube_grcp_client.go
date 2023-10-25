package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "p4kube/p4kube_grpc"
	v1alpha1 "p4kube/pkg/apis/p4kube/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func checkStatus(finished chan string, stream pb.P4DeploymentService_DeployP4Client) {
	var deploymentStatus *pb.DeploymentStatus
	log.Print(deploymentStatus)
	for {
		deploymentStatus, _ = stream.Recv()
		log.Printf("Status: %v\n", deploymentStatus)
		if deploymentStatus != nil {
			if deploymentStatus.Status == "Deployed" {
				log.Println("P4 Resource successfully deployed")
				break
			}
			if deploymentStatus.Status == "UnDeployed" {
				log.Printf("Could not deploy P4 resource")
				break
			}
		} else {
			log.Println("Status: Deploying")
			time.Sleep(time.Second)
			continue
		}
	}
	finished <- deploymentStatus.Status
}

func createCommand(p4resource *v1alpha1.P4) string {
	cmdExec := fmt.Sprintf("cd %v; %v %v", p4resource.Spec.CompilerDirectory, p4resource.Spec.CompilerCommand, p4resource.Spec.P4Program)

	return cmdExec
}

func dialWorkerNode(p4resource *v1alpha1.P4, address string) string {
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
		log.Fatalf("Could not receive response from worker node: %v", err)
	}

	finished := make(chan string)

	go checkStatus(finished, stream)

	deploymentStatus := <-finished
	return deploymentStatus
}
