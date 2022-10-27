package repo

import (
	"game-score/internal/savescore"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type SaveScoreRepository struct {
	ddb dynamodbiface.DynamoDBAPI
}

func NewSaveScoreRepository(ddb dynamodbiface.DynamoDBAPI) *SaveScoreRepository {
	return &SaveScoreRepository{
		ddb: ddb,
	}
}

func (r *SaveScoreRepository) SaveScore(score savescore.Score) error {
	tableName := "game-score-table"

	scoreMap, err := dynamodbattribute.MarshalMap(score)
	if err != nil {
		return err
	}

	_, err = r.ddb.PutItem(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      scoreMap,
	})

	return err
}
