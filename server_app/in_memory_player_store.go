package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{make(map[string]int), sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	scores map[string]int
	sync.Mutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score, ok := i.scores[name]
	if !ok {
		return -1
	}
	return score
}

func (i *InMemoryPlayerStore) IncrementPlayerScore(name string) {
	i.scores[name]++
}
