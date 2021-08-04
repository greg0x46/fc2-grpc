package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/greg0x46/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main () {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if(err != nil){
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if(err != nil){
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient){
	req := &pb.User {
		Id: "0",
		Name: "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)
	if(err != nil){
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		
		if(err == io.EOF){
			break
		}
		
		if(err != nil){
			log.Fatalf("Could not receive the msg from stram: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}