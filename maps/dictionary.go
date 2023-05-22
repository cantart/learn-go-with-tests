package dictionary

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type Dictionary map[string]string

// create DictionaryErr type which implements the error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	// The second value is a boolean which indicates if the key was found successfully.
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition

	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

/*
In Go, when you declare a map variable without initializing it, such as `var m map[string]string`,
the variable `m` is initialized with the zero value for a map, which is `nil`.
This means that `m` does not refer to a valid map and does not have an associated memory allocation.

When you initialize an empty map using `var dictionary = map[string]string{}`,
it creates a new map and assigns it to the `dictionary` variable.
In this case, `dictionary` is a valid map with an associated memory allocation,
and it does not have a `nil` value. Therefore, accessing or modifying `dictionary`
will not lead to a runtime panic.

To summarize:

* `var m map[string]string` initializes `m` as a map variable with the zero value, which is `nil`. It does not have an associated memory allocation for the map.
* `var dictionary = map[string]string{}` initializes `dictionary` as a map variable and allocates memory for an empty map. It does not have a `nil` value and is ready to be used without causing a runtime panic.

By properly initializing the map using either a composite literal or the `make` function,
you ensure that the variable refers to a valid map with allocated memory,
avoiding runtime errors when interacting with the map.
*/
