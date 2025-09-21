package api

import (
	"context"
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUser(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(event.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, err
	}


	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Username and password are required",
		}, nil
	}

	ctx := context.Background()

	userExists, err := api.dbStore.DoesUserExist(ctx, registerUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, err
	}

	if userExists {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusConflict,
			Body:       "User already exists",
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}, fmt.Errorf("failed to register user")
}
