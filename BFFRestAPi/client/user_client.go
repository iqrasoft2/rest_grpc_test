package client

import (
	"context"
	"errors"

	"rest_grpc_test/model/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Users struct {
	Description     string `json:"Description"`
}

type UsersClient struct {
}

var (
	userServiceAddress = "localhost:5001"
	userServiceClient user.UserServiceClient
)

func prepareAdviceGrpcClient(c *context.Context) error {

	conn, err := grpc.DialContext(*c, userServiceAddress, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		userServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if userServiceClient != nil {
		conn.Close()
		return nil
	}

	userServiceClient = user.NewUserServiceClient(conn)
	return nil
}

func (ac *UsersClient) GreetUser(c *context.Context) (*Users, error) {

	if err := prepareAdviceGrpcClient(c); err != nil {
		return nil, err
	}

	response, err := userServiceClient.GreetUser(*c, &user.GreetingRequest{
		Name: "Hallo Sobat IT CRM",
		Salutation: "- Rest to gRpc -",
	})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}

	return &Users{Description: response.GreetingMessage}, nil
}
