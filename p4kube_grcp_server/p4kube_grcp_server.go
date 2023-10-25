package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	pb "p4kube/p4kube_grpc"

	"google.golang.org/grpc"
)

// server port where service is being hosted
const (
	port = ":50051"
)

type DeployP4Server struct {
	pb.UnimplementedP4DeploymentServiceServer
}

func (s *DeployP4Server) DeployP4(in *pb.P4Resource, stream pb.P4DeploymentService_DeployP4Server) error {
	status := "Deploying"

	log.Printf("Received P4Resource: %v\n", in.GetName())
	cmd := exec.Command("/bin/sh", "-c", in.Command)
	log.Print("Command to be executed: ", cmd, "\n")

	// Print output of command, also the error if command not successful.
	log.Println("Showing the logs of deploying P4 resource with t4p4s")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		status = "UnDeployed"
		log.Fatalf("Deployment unsuccessful: %v", err)
	} else {
		status = "Deployed"
		log.Printf("Deployment successful")
	}

	for {
		if status == "Deployed" {
			break
		}

		log.Printf("Sending status: %v to client", status)

		err := stream.Send(&pb.DeploymentStatus{Status: status})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
		time.Sleep(time.Second)
	}

	log.Printf("Status: %s", status)

	err = stream.Send(&pb.DeploymentStatus{Status: status})
	if err != nil {
		log.Fatalf("Error while sending data to client: %v", err)
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed listening to server port: %v", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterP4DeploymentServiceServer(server, &DeployP4Server{})

	log.Printf("Server listening at %v", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve at server address: %v", err.Error())
	}
}
