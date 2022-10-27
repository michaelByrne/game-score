package repo

import (
	"game-score/internal/savescore"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type fakeScoreSaver struct {
	dynamodbiface.DynamoDBAPI
	called     bool
	calledWith *dynamodb.PutItemInput
}

func (f *fakeScoreSaver) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	f.called = true
	f.calledWith = input
	return nil, nil
}

func TestSaveScoreRepository_SaveScore(t *testing.T) {
	ddb := &fakeScoreSaver{}
	repo := NewSaveScoreRepository(ddb)

	score := savescore.Score{
		ID:     "fake-id",
		Team:   "fake-team",
		Points: 100,
		Home:   true,
	}

	err := repo.SaveScore(score)
	require.NoError(t, err)

	require.True(t, ddb.called)
	assert.Equal(t, "fake-id", *ddb.calledWith.Item["scoreId"].S)
	assert.Equal(t, "fake-team", *ddb.calledWith.Item["team"].S)
	assert.Equal(t, "100", *ddb.calledWith.Item["points"].N)
	assert.Equal(t, true, *ddb.calledWith.Item["home"].BOOL)
}
