package repo

import (
	"game-score/internal/getscore"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type GetScoreRepository struct {
	ddb dynamodbiface.DynamoDBAPI
}

func NewGetScoreRepository(ddb dynamodbiface.DynamoDBAPI) *GetScoreRepository {
	return &GetScoreRepository{
		ddb: ddb,
	}
}

func (r *GetScoreRepository) GetScore(id string) (*getscore.Score, error) {
	tableName := "game-score-table"

	result, err := r.ddb.GetItem(&dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"scoreId": {
				S: &id,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var score getscore.Score
	if err := dynamodbattribute.UnmarshalMap(result.Item, &score); err != nil {
		return nil, err
	}

	return &score, nil
}
