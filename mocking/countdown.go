package mocking

import (
  "fmt"
  "io"
  "os"
  "time"
)

func main() {
  sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
  Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
  Sleep()
}

type ConfigurableSleeper struct {
 duration time.Duration
 sleep func(duration time.Duration)
}

func (s *ConfigurableSleeper) Sleep(){
  s.sleep(s.duration)
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