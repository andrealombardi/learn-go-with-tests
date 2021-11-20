package propbasedtests

import "testing"

func TestRomanRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		ArabicValue int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range cases {
		got := ConvertToRoman(test.ArabicValue)
		if got != test.Want {
			t.Errorf("want %q, got %q", test.Want, got)
		}
	}
}
