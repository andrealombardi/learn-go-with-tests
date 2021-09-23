package arrslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Test with slice", func(t *testing.T) {
		want := 6
		integers := []int{1, 2, 3}
		got := Sum(integers)

		if want != got {
			t.Errorf("Wanted %d, got %d, given %v", want, got, integers)
		}
	})
}

func TestSumAll(t *testing.T) {
	want := []int{6, 3}
	got := SumAll([]int{1, 2, 3}, []int{1, 2})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T){

	checksum := func(tb testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got){
			t.Errorf("Wanted %v, got %v", want, got)
		}
	}
	t.Run("All supplied slices are populated", func(t *testing.T) {
		want := []int {5,2}
		got := SumAllTails([]int{1,2,3}, []int{1,2})
		checksum(t, want, got)
	})

	t.Run("One element slices evaluate to zero", func(t *testing.T) {
		want := []int {5,0}
		got := SumAllTails([]int{1,2,3}, []int{2})
		checksum(t, want, got)
	})

	t.Run("Empty slices evaluate to zero", func(t *testing.T) {
		want := []int {5,0}
		got := SumAllTails([]int{1,2,3}, []int{})
		checksum(t, want, got)
	})

}