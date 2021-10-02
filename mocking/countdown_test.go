package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Test Output", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if 4 != sleeper.Calls {
			t.Errorf("got %d calls want %d", 4, sleeper.Calls)
		}
	})

	t.Run("Test sleeping pattern", func(t *testing.T) {
		spySleepWriter := &SpyCountdownOperations{}
		Countdown(spySleepWriter, spySleepWriter)

		want := []string {
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepWriter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepWriter.Calls)
		}


	})
}

func TestConfigurableSleeper_Sleep(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slpet for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}