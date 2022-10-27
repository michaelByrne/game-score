package main

import (
	"game-score/cmd/app"
	"game-score/internal/getscore"
	getrepo "game-score/internal/getscore/repo"
	"game-score/internal/savescore"
	saverepo "game-score/internal/savescore/repo"
	"game-score/internal/transport"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ddb := dynamodb.New(sess)

	getScoreRepo := getrepo.NewGetScoreRepository(ddb)
	getScore := getscore.NewGetScore(getScoreRepo)

	saveScoreRepo := saverepo.NewSaveScoreRepository(ddb)
	saveScore := savescore.NewSaveScore(saveScoreRepo)

	handlers := transport.NewHandlers(getScore, saveScore)

	a := app.NewApp(handlers)

	log.Fatal(a.Run())
}
