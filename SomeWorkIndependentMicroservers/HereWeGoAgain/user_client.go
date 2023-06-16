package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "HereWeGoAgain/user" // Import the generated protobuf code
)

func main() {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a UserService client
	client := pb.NewUserServiceClient(conn)

	// Send a registration request
	registrationRequest := &pb.RegistrationRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	registeredUser, err := client.Register(context.Background(), registrationRequest)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	log.Println("Registered User:")
	log.Println(registeredUser)

	// Send an authorization request
	authorizationRequest := &pb.AuthorizationRequest{
		Email:    "john@example.com",
		Password: "password123",
	}
	authorizedUser, err := client.Authorize(context.Background(), authorizationRequest)
	if err != nil {
		log.Fatalf("Authorization failed: %v", err)
	}
	log.Println("Authorized User:")
	log.Println(authorizedUser)
}
