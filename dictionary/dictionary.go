package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Could not find the word searched for.")

func (d Dictionary) Search(term string) (string, error) {
	definition, found := d[term]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}
