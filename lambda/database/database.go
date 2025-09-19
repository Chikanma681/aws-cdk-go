package database

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
    databaseStore *dynamodb.Client
}

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