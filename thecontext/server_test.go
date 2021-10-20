package thecontext

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("request, response ok", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&StubStore{data, false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

	})

	t.Run("tells store to cancel workif request is cancelled", func(t *testing.T){

	})

}
