package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("Could not find the word searched for.")
	ErrAlreadyExists    = DictionaryErr("Word already exists.")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, exists := d[term]
	if exists {
		return ErrAlreadyExists
	}

	d[term] = definition
	return nil
}

func (d Dictionary) Update(term, definition string) error {
	_, exists := d[term]
	if !exists {
		return ErrWordDoesNotExist
	}

	d[term] = definition
	return nil
}
