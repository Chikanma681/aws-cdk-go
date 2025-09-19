package main

import (
    "context"
    "fmt"
    "lambda-func/app"
    "log"
    "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
    Username string `json:"username"`
}

var application app.App

// Take in a payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
    if event.Username == "" {
        return "", fmt.Errorf("username is required")
    }
    return fmt.Sprintf("Successfully called by - %s!", event.Username), nil
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