package main

import (
	"context"
	"fmt"
	"lambda-func/app"
	"lambda-func/types"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var application app.App

// Take in a payload and do something with it
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is required")
	}
	if event.Password == "" {
		return "", fmt.Errorf("password is required")
	}

	// Create RegisterUser object
	registerUser := types.RegisterUser{
		Username: event.Username,
		Password: event.Password,
	}

	// Register the user using the API handler
	err := application.ApiHandler.RegisterUser(ctx, registerUser)
	if err != nil {
		return "", fmt.Errorf("registration failed: %w", err)
	}

	return fmt.Sprintf("Successfully registered user - %s!", event.Username), nil
}

func main() {
	ctx := context.Background()
	var err error
	application, err = app.NewApp(ctx)
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}

	lambda.Start(HandleRequest)
}
