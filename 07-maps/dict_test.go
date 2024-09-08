package dictionary

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Expected an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertDefinition := func(t testing.TB, dictionary Dictionary, word, definition string) {
		t.Helper()

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, definition)
	}
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, "this is just a test")
	})
	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("hide")
		assertError(t, err, ErrNotFound)
		assertStrings(t, got, "")
	})
	t.Run("add", func(t *testing.T) {
		word := "testa"
		definition := "just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
}
