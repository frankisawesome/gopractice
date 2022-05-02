package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

type SpyPlayerStore struct {
	scores map[string]int
	sync.Mutex
}

func (s *SpyPlayerStore) GetPlayerScore(name string) int {
	s.Lock()
	score, ok := s.scores[name]
	s.Unlock()
	if !ok {
		return -1
	}
	return score
}

func (s *SpyPlayerStore) IncrementPlayerScore(player string) {
	s.Lock()
	s.scores[player]++
	s.Unlock()
}

func TestGETPlayers(t *testing.T) {
	store := SpyPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  5,
		},
		sync.Mutex{},
	}
	server := &PlayerServer{&store}

	t.Run("return Pepper's score", func(t *testing.T) {
		got := getPlayerResponse("Pepper", server)
		assertResponseBody(t, got, "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		got := getPlayerResponse("Floyd", server)
		assertResponseBody(t, got, "5")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		got := getPlayerResponse("Apollo", server)

		assertResponseStatus(t, got, http.StatusNotFound)
		assertResponseBody(t, got, "Not found")
	})
}

func TestStoreWins(t *testing.T) {
	store := SpyPlayerStore{
		scores: map[string]int{
			"Pepper": 0,
		},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted and score on POST", func(t *testing.T) {
		response := postPlayerResponse("Pepper", server)

		assertResponseStatus(t, response, http.StatusAccepted)
		assertResponseBody(t, response, "1")
	})

	t.Run("it increments stored score on POST", func(t *testing.T) {
		currentScoreStr := getPlayerResponse("Pepper", server).Body.String()
		currentScore, _ := strconv.Atoi(currentScoreStr)

		postPlayerResponse("Pepper", server)

		newScoreStr := getPlayerResponse("Pepper", server).Body.String()
		newScore, _ := strconv.Atoi(newScoreStr)

		if newScore != currentScore+1 {
			t.Errorf("expected score to be %d, got %d", currentScore+1, newScore)
		}
	})

	t.Run("it returns not found if player doesn't exist", func(t *testing.T) {
		response := postPlayerResponse("Fake", server)

		assertResponseStatus(t, response, http.StatusNotFound)
		assertResponseBody(t, response, "Not found")
	})
}

func TestConcurrency(t *testing.T) {
	store := SpyPlayerStore{
		scores: map[string]int{
			"Pepper": 0,
		},
	}
	server := &PlayerServer{&store}

	t.Run("concurrent writes should work", func(t *testing.T) {
		const score = 1000
		var wg sync.WaitGroup
		wg.Add(score)
		for i := 0; i < score; i++ {
			go func() {
				defer wg.Done()
				postPlayerResponse("Pepper", server)
			}()
		}
		wg.Wait()

		response := getPlayerResponse("Pepper", server)
		assertResponseBody(t, response, "1000")
	})
}

func assertResponseBody(t testing.TB, got *httptest.ResponseRecorder, want string) {
	t.Helper()
	if got.Body.String() != want {
		t.Errorf("response body is wrong, got %q want %q", got.Body.String(), want)
	}
}

func assertResponseStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("response status code is wrong, got %d want %d", got.Code, want)
	}
}

func getPlayerResponse(name string, server *PlayerServer) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	return response
}

func postPlayerResponse(name string, server *PlayerServer) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	return response
}
