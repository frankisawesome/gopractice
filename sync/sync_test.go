package sync_prac

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)

	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, *counter, wantedCount)
	})
}

func assertCounter(t *testing.T, c Counter, n int) {
	t.Helper()
	if c.Value() != n {
		t.Errorf("got %d, want %d", c.Value(), n)
	}
}
