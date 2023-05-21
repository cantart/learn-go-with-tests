package dictionary

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	// The second value is a boolean which indicates if the key was found successfully.
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
