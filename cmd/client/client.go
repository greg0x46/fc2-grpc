package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
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
	//AddUserVerbose(client)
	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient){
	reqs := []*pb.User{
		&pb.User{
			Id: "1",
			Name: "Greg1",
			Email: "greg@greg.com",
		},
		&pb.User{
			Id: "2",
			Name: "Greg2",
			Email: "greg@greg.com",
		},
		&pb.User{
			Id: "3",
			Name: "Greg3",
			Email: "greg@greg.com",
		},
		&pb.User{
			Id: "4",
			Name: "Greg4",
			Email: "greg@greg.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if(err != nil){
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}