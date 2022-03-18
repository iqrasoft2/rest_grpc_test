package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rest_grpc_test/model/user"
	"log"
	"time"
)

const userServiceAddress = "localhost:5001"

func main(){
	//Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	//Create new userService client
	userServiceClient := user.NewUserServiceClient(conn)

	//create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	response, err := userServiceClient.GreetUser(ctx, &user.GreetingRequest{
		Name: "Hallo Sobat IT CRM",
		Salutation: "- gRPC to gRPC -",
	})
	if err != nil{
		log.Fatalf("Could not create request: %v", err)
	}
  
  //show response
	fmt.Println(response.GreetingMessage)
}