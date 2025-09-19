package app

import (
	"lambda-func/api"
	"lambda-func/database"
	"context"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp(ctx context.Context) (App ,error) {
	db, err := database.NewDynamoDBClient(ctx);
	
	if err != nil {
		return App{}, err
	}

	apiHandler := api.NewApiHandler(db);
	return App{
		ApiHandler: apiHandler,
	}, nil
}