package getscore

import "time"

type ScoreGetter interface {
	GetScore(id string) (*Score, error)
}

type Score struct {
	ID     string     `json:"scoreId"`
	Team   string     `json:"team"`
	Points int        `json:"points"`
	Date   *time.Time `json:"date"`
	Home   bool       `json:"home"`
}

type GetScore struct {
	scoreGetter ScoreGetter
}

func NewGetScore(scoreGetter ScoreGetter) *GetScore {
	return &GetScore{
		scoreGetter: scoreGetter,
	}
}

func (g *GetScore) GetScore(id string) (*Score, error) {
	return g.scoreGetter.GetScore(id)
}
