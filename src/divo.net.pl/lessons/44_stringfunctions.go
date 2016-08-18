// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.
package lessons

import(
	"fmt"
	str "strings"
)

func StringFunctions() {
	// We alias `fmt.Println` to a shorter name as we'll use
	// it a lot below.
	var p = fmt.Println
	
	// Here's a sample of the functions available in
	// `strings`. Note that these are all functions from
	// package, not methods on the string object itself.
	// This means that we need pass the string in question
	// as the first argument to the function.
	p("Contains:  ", str.Contains("test", "es"))
	p("Count:     ", str.Count("test", "t"))
	p("HasPrefix: ", str.HasPrefix("test", "te"))
	p("HasSuffix: ", str.HasSuffix("test", "st"))
	p("Index:     ", str.Index("test", "e"))
	p("Join:      ", str.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", str.Repeat("a", 5))
	p("Replace:   ", str.Replace("foo", "o", "0", -1))
	p("Replace:   ", str.Replace("foo", "o", "0", 1))
	p("Split:     ", str.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", str.ToLower("TEST"))
	p("ToUpper:   ", str.ToUpper("test"))
	p()
	
	// You can find more functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.

	// Not part of `strings` but worth mentioning here are
	// the mechanisms for getting the length of a string
	// and getting a character by index.
	p("Len:       ", len("hello"))
	p("Char:      ", "hello"[1])
}