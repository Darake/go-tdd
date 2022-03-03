package main

type Dictionary map[string]string

func (d Dictionary) Search(term string) string {
	return d[term]
}
