package mocking

import (
  "fmt"
  "io"
  "os"
  "time"
)

func main() {
  Countdown(os.Stdout, &DefaultSleeper{})
}

type Sleeper interface {
  Sleep()
}

type DefaultSleeper struct {
}

func (s *DefaultSleeper) Sleep() {
 time.Sleep(1 * time.Second)
}

const countdownStart = 3
const finalWord = "Go!"

func Countdown(out io.Writer, s Sleeper) {
  for i := countdownStart; i > 0; i-- {
    s.Sleep()
    fmt.Fprintln(out, i)
  }
  s.Sleep()
  fmt.Fprintf(out, finalWord)
}