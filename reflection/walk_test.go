package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{"Struct with a string and an int",
			struct {
				Name string
				Age  int
			}{"Chris", 44},
			[]string{"Chris"},
		},
		{"Non flat structure",
			Person{
				"Chris",
				Profile{
					"London",
					45,
				}},
			[]string{"Chris", "London"},
		},
		{"Pointer to things",
			&Person{"Christos",
				Profile{"London", 45}},
			[]string{"Christos", "London"},
		},
		{
			"Slices",
			[]Profile {
				{"London", 44},
				{"Brescia", 23},
			},
			[]string{"London", "Brescia"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if len(got) != len(test.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got %d want %d", len(got), len(test.ExpectedCalls))
			}
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}

}
