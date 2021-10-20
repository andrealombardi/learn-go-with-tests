package synccounter

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, wantedCount)

	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

type T struct {
	Version int `json:"version"`
	Offers  struct {
		Product []struct {
			Genre    string   `json:"genre"`
			Director string   `json:"director"`
			Title    string   `json:"title"`
			Year     int      `json:"year"`
			Staring  []string `json:"staring,omitempty"`
		} `json:"product"`
		Bakset []struct {
			Genre    string   `json:"genre"`
			Director string   `json:"director"`
			Title    string   `json:"title"`
			Year     int      `json:"year"`
			Staring  []string `json:"staring,omitempty"`
		} `json:"bakset"`
		Metadata struct {
			LastUpdated time.Time `json:"lastUpdated"`
		} `json:"metadata"`
		StringJson string `json:"stringJson"`
	} `json:"offers"`
}