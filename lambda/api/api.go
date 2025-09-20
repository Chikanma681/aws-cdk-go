package api

import (
	"context"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUser(ctx context.Context, event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	userExists, err := api.dbStore.DoesUserExist(ctx, event.Username)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}

	if userExists {
		return fmt.Errorf("user with username %s already exists", event.Username)
	}

	err = api.dbStore.InsertUser(ctx, event)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
