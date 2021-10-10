package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout of %v exceeded while waiting for %v and %v", timeout, a, b)
	}
}

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func() {
		http.Get(url)
		close(c)
	}()
	return c
}
