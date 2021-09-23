package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}
}

func ExampleRepeat() {
	got := Repeat("b", 5)
	fmt.Print(got)
	//Output: bbbbb
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
