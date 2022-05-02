package main

import (
	"fmt"
	"net/http"
	"strings"
)

const PlayerNotFoundScore = -1

type PlayerStore interface {
	GetPlayerScore(name string) int
	IncrementPlayerScore(name string)
	// Mutex
	Lock()
	Unlock()
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if p.store.GetPlayerScore(player) == PlayerNotFoundScore {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found")
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, p.store.GetPlayerScore(player))
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		p.store.IncrementPlayerScore(player)
		fmt.Fprint(w, p.store.GetPlayerScore(player))
	}
}
