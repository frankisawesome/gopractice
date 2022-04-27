package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word")
var ErrWordExists = errors.New("word already exists")

func (d Dictionary) Search(w string) (string, error) {
	value, ok := d[w]
	if ok {
		return value, nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(k, v string) error {
	_, ok := d[k]
	if !ok {
		d[k] = v
		return nil
	} else {
		return ErrWordExists
	}
}

func Search(d map[string]string, w string) (value string) {
	return d[w]
}
