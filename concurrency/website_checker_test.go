package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "this://we.want.to.fail" {
		return false
	}
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://amazon.com",
		"this://we.want.to.fail",
	}
	want := map[string]bool{
		"http://google.com": true,
		"http://amazon.com": true,
		"this://we.want.to.fail": false,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}