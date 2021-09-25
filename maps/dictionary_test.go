package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Insert new", func(t *testing.T) {

		dictionary := Dictionary{}
		definition := "this is just a test"
		word := "test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Insert existing", func(t *testing.T) {

		definition := "this is just a test"
		word := "test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "this is another definition")

		assertError(t, err, ErrAlreadyExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Update existing", func(t *testing.T) {

		definition := "this is just a test"
		word := "test"
		updated := "this is just a test - updated"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, updated)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updated)
	})

}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if got != definition {
		t.Errorf("got %q, definition %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
