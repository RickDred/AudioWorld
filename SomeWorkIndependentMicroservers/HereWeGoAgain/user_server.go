package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "HereWeGoAgain/user"
)

func validateEmail(email string) bool {
	// Email validation regular expression pattern
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func validateName(name string) bool {
	// Name should not be empty and should not contain any digits
	return len(strings.TrimSpace(name)) > 0 && !hasDigits(name)
}

func hasDigits(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func validatePassword(password string) bool {
	// Password should have at least 8 characters
	return len(password) >= 8
}

func Validate(email, name, password string) (err bool, message string) {

	// Validate email
	if !validateEmail(email) {
		return true, "Invalid Email"
	}

	// Validate name
	if !validateName(name) {
		return true, "Invalid name"
	}

	// Validate password
	if !validatePassword(password) {
		return true, "Invalid password"
	}

	return false, "User valid"
}

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegistrationRequest) (*pb.User, error) {
	err, msg := Validate(req.GetEmail(), req.GetName(), req.GetPassword())
	if err {
		log.Fatalf("%v", msg)
	}
	// There should be adding out user to db
	user := &pb.User{
		Id:       uuid.New().String(),
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	return user, nil
}

func (s *UserServiceServer) Authorize(ctx context.Context, req *pb.AuthorizationRequest) (*pb.User, error) {
	// getting user from non existing db
	authorizedUser := &pb.User{
		Id:       uuid.New().String(),
		Name:     "John Doe", // we should take name from db, but .....
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	return authorizedUser, nil
}

func main() {
	// Create a gRPC server
	server := grpc.NewServer()

	userService := &UserServiceServer{}
	pb.RegisterUserServiceServer(server, userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server started on port 50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
