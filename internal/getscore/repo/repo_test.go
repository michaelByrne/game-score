package repo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type fakeDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	called     bool
	calledWith *dynamodb.GetItemInput
	response   *dynamodb.GetItemOutput
}

func (f *fakeDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	f.called = true
	f.calledWith = input
	return f.response, nil
}

func TestGetScoreRepository_GetScore(t *testing.T) {
	scoreID := "fake-id"

	ddb := &fakeDynamoDB{
		response: &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"scoreId": {
					S: &scoreID,
				},
			},
		},
	}

	repo := NewGetScoreRepository(ddb)

	score, err := repo.GetScore(scoreID)
	require.NoError(t, err)

	require.True(t, ddb.called)
	assert.Equal(t, scoreID, *ddb.calledWith.Key["scoreId"].S)
	assert.Equal(t, scoreID, score.ID)
}
