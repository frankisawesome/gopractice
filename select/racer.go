package racer

import (
	"net/http"
	"time"
)

type Results map[string]float64

func Racer(url1, url2 string) (winner string, results Results) {
	//time1 := measureResponseTime(url1)
	//time2 := measureResponseTime(url2)
	//
	//results = Results{url1: time1, url2: time2}
	//if time1 > time2 {
	//	winner = url2
	//} else {
	//	winner = url1
	//}
	//return
	select {
	case <-ping(url1):
		return url1, Results{}
	case <-ping(url2):
		return url2, Results{}
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) (resultTime float64) {
	start := time.Now()
	http.Get(url)
	resultTime = time.Since(start).Seconds()
	return
}
