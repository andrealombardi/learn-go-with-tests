package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Select the fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(0 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got, _ := Racer(slowUrl, fastUrl, 3*time.Second)
		want := fastUrl

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})

	t.Run("Should timeout after 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(1 * time.Second)
		serverB := makeDelayedServer(1 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 200*time.Millisecond)

		if err == nil {
			t.Error("Expected timeout error")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
