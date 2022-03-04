package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	existingTerm := "existing"
	originalDefinition := "og"
	dictionary := Dictionary{existingTerm: originalDefinition}

	t.Run("new word", func(t *testing.T) {
		term := "test"
		definition := "new"

		err := dictionary.Add(term, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add(existingTerm, "new definition")

		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, existingTerm, originalDefinition)
	})
}

func TestUpdate(t *testing.T) {
	existingTerm := "existing"
	dictionary := Dictionary{existingTerm: "original"}

	t.Run("existing word", func(t *testing.T) {
		newWord := "new"
		err := dictionary.Update(existingTerm, newWord)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, existingTerm, newWord)
	})

	t.Run("new word", func(t *testing.T) {
		newWord := "new"
		err := dictionary.Update("newTerm", newWord)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	existingTerm := "existing"
	dictionary := Dictionary{existingTerm: "word"}

	dictionary.Delete(existingTerm)

	_, err := dictionary.Search(existingTerm)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", existingTerm)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, term, definition string) {
	t.Helper()
	got, err := dictionary.Search(term)
	assertError(t, err, nil)
	assertStrings(t, got, definition)
}
