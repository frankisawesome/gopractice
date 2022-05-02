package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{
		map[string]int{
			"Pepper": 0,
		},
		sync.Mutex{},
	}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
