package transport

import (
	"encoding/json"
	"game-score/internal/getscore"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type GetScore interface {
	GetScore(id string) (*getscore.Score, error)
}

type SaveScore interface {
	SaveScore(team string, points int, date *time.Time, home bool) error
}

type Handlers struct {
	getScore  GetScore
	saveScore SaveScore
}

func NewHandlers(getScore GetScore, saveScore SaveScore) *Handlers {
	return &Handlers{
		getScore:  getScore,
		saveScore: saveScore,
	}
}

func (h *Handlers) GetScoreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	score, err := h.getScore.GetScore(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if score == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("score not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(score)
}

func (h *Handlers) SaveScoreHandler(w http.ResponseWriter, r *http.Request) {
	var score getscore.Score
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.saveScore.SaveScore(score.Team, score.Points, score.Date, score.Home)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
