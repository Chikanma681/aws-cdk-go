package database

import (
	"context"
	lambda_types "lambda-func/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.Client
}

const (
	TABLE_NAME = "userTable"
)

func NewDynamoDBClient(ctx context.Context) (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	db := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{
		databaseStore: db,
	}, nil
}

func (u DynamoDBClient) DoesUserExist(ctx context.Context, username string) (bool, error) {
	result, err := u.databaseStore.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: username},
		},
		TableName: aws.String(TABLE_NAME),
	})
	return err == nil && result.Item != nil, err
}

func (u *DynamoDBClient) InsertUser(ctx context.Context, user lambda_types.RegisterUser) error {
	// Assemble the item using correct AWS SDK v2 syntax
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: user.Username},
			"password": &types.AttributeValueMemberS{Value: user.Password},
		},
	}

	// Insert the item with context
	_, err := u.databaseStore.PutItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
