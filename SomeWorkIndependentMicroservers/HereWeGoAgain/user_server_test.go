package main

import (
	"context"
	"testing"

	pb "HereWeGoAgain/user"
)

func TestValidateEmail(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "invalid-email"

	if !validateEmail(validEmail) {
		t.Errorf("Expected email validation to pass for %s, but it failed", validEmail)
	}

	if validateEmail(invalidEmail) {
		t.Errorf("Expected email validation to fail for %s, but it passed", invalidEmail)
	}
}

func TestValidateName(t *testing.T) {
	validName := "John Doe"
	invalidName := "John Doe 123"

	if !validateName(validName) {
		t.Errorf("Expected name validation to pass for %s, but it failed", validName)
	}

	if validateName(invalidName) {
		t.Errorf("Expected name validation to fail for %s, but it passed", invalidName)
	}
}

func TestValidatePassword(t *testing.T) {
	validPassword := "password123"
	invalidPassword := "pass"

	if !validatePassword(validPassword) {
		t.Errorf("Expected password validation to pass for %s, but it failed", validPassword)
	}

	if validatePassword(invalidPassword) {
		t.Errorf("Expected password validation to fail for %s, but it passed", invalidPassword)
	}
}

func TestValidate(t *testing.T) {
	validEmail := "test@example.com"
	invalidEmail := "invalid-email"
	validName := "John Doe"
	invalidName := "John Doe 123"
	validPassword := "password123"
	invalidPassword := "pass"

	err, msg := Validate(validEmail, validName, validPassword)
	if err {
		t.Errorf("Expected user validation to pass, but it failed with error: %s", msg)
	}

	err, msg = Validate(invalidEmail, validName, validPassword)
	if !err {
		t.Errorf("Expected user validation to fail for invalid email, but it passed")
	}

	err, msg = Validate(validEmail, invalidName, validPassword)
	if !err {
		t.Errorf("Expected user validation to fail for invalid name, but it passed")
	}

	err, msg = Validate(validEmail, validName, invalidPassword)
	if !err {
		t.Errorf("Expected user validation to fail for invalid password, but it passed")
	}
}

func TestUserServiceServer_Register(t *testing.T) {
	// Create a UserServiceServer instance
	userService := &UserServiceServer{}

	// Create a RegistrationRequest for testing
	req := &pb.RegistrationRequest{
		Email:    "test@example.com",
		Name:     "John Doe",
		Password: "password123",
	}

	// Create a context for the RPC call
	ctx := context.Background()

	// Call the Register method and check the response
	res, err := userService.Register(ctx, req)
	if err != nil {
		t.Errorf("Register failed with error: %v", err)
	}

	// Verify the returned user
	if res == nil || res.Id == "" || res.Name != req.Name || res.Email != req.Email || res.Password != req.Password {
		t.Errorf("Register returned an invalid user")
	}
}

func TestUserServiceServer_Authorize(t *testing.T) {
	// Create a UserServiceServer instance
	userService := &UserServiceServer{}

	// Create an AuthorizationRequest for testing
	req := &pb.AuthorizationRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Create a context for the RPC call
	ctx := context.Background()

	// Call the Authorize method and check the response
	res, err := userService.Authorize(ctx, req)
	if err != nil {
		t.Errorf("Authorize failed with error: %v", err)
	}

	// Verify the returned user
	if res == nil || res.Id == "" || res.Name != "John Doe" || res.Email != req.Email || res.Password != req.Password {
		t.Errorf("Authorize returned an invalid user")
	}
}
