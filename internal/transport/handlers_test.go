package transport

import (
	"game-score/internal/getscore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const testResponse = `{"scoreId":"test-id","team":"blazers","points":0,"date":null,"home":false}
`

type fakeScoreSaver struct {
	called         bool
	calledWithTeam string
}

func (f *fakeScoreSaver) SaveScore(team string, points int, date *time.Time, home bool) error {
	f.called = true
	f.calledWithTeam = team
	return nil
}

type fakeScoreFetcher struct {
	called     bool
	calledWith string
	response   *getscore.Score
}

func (f *fakeScoreFetcher) GetScore(id string) (*getscore.Score, error) {
	f.called = true
	f.calledWith = id
	return f.response, nil
}

func TestHandlers_GetScoreHandler(t *testing.T) {
	t.Run("it returns a 404 if the score is not found", func(t *testing.T) {
		fetcher := &fakeScoreFetcher{}
		handler := NewHandlers(fetcher, nil)

		req, err := http.NewRequest("GET", "/scores/unknown", nil)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		h := http.HandlerFunc(handler.GetScoreHandler)

		h.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.Equal(t, "score not found", recorder.Body.String())
	})

	t.Run("it returns a 200 if the score is found", func(t *testing.T) {
		fetcher := &fakeScoreFetcher{
			response: &getscore.Score{
				ID:   "test-id",
				Team: "blazers",
			},
		}
		handler := NewHandlers(fetcher, nil)

		req, err := http.NewRequest("GET", "/scores/test-id", nil)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		h := http.HandlerFunc(handler.GetScoreHandler)

		h.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, testResponse, recorder.Body.String())
	})
}

func TestHandlers_SaveScoreHandler(t *testing.T) {
	t.Run("it returns a 200 if the score is saved", func(t *testing.T) {
		saver := &fakeScoreSaver{}
		handler := NewHandlers(nil, saver)

		body := `{"team":"blazers","points":0,"date":null,"home":false}`

		req, err := http.NewRequest("POST", "/scores", strings.NewReader(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		h := http.HandlerFunc(handler.SaveScoreHandler)

		h.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "blazers", saver.calledWithTeam)
	})
}
