package reflection

import (
	"reflect"
)

/*
! About Reflection
* Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.
* You can't use NumField on a pointer Value, we need to extract the underlying value before we can do that by using Elem().
*/

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// When you're doing a comparison on the same value more than once generally refactoring into a switch will improve readability and make your code easier to extend.
	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
