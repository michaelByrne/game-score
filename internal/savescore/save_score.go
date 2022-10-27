package savescore

import (
	"github.com/google/uuid"
	"time"
)

type ScoreSaver interface {
	SaveScore(score Score) error
}

type Score struct {
	ID     string     `json:"scoreId"`
	Team   string     `json:"team"`
	Points int        `json:"points"`
	Date   *time.Time `json:"date"`
	Home   bool       `json:"home"`
}

type SaveScore struct {
	scoreSaver ScoreSaver
}

func NewSaveScore(scoreSaver ScoreSaver) *SaveScore {
	return &SaveScore{
		scoreSaver: scoreSaver,
	}
}

func (s *SaveScore) SaveScore(team string, points int, date *time.Time, home bool) error {
	score := Score{
		ID:     uuid.New().String(),
		Team:   team,
		Points: points,
		Date:   date,
		Home:   home,
	}

	return s.scoreSaver.SaveScore(score)
}
