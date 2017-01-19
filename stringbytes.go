// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import  ("strings"
	"fmt"
	"bytes"
)

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.


func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.


	fmt.Println("Contains:  ", bytes.Contains([]byte("test"), []byte("es")))
	fmt.Println("Count:     ", bytes.Count([]byte("test"), []byte("t")))
	fmt.Println("HasPrefix: ", bytes.HasPrefix([]byte("test"), []byte("te")))
	fmt.Println("HasSuffix: ", bytes.HasSuffix([]byte("test"), []byte("st")))
	fmt.Println("Index:     ", bytes.Index([]byte("test"), []byte("e")))
	fmt.Println("Replace:   ", string(bytes.Replace([]byte("foo"), []byte("o"), []byte("0"), -1)))
	fmt.Println("Replace:   ", string(bytes.Replace([]byte("foo"), []byte("o"), []byte("0"), 1)))
	fmt.Println("Split:     ", bytes.Split([]byte("a-b-c-d-e"), []byte("-")))
	fmt.Println("ToLower:   ", string(bytes.ToLower([]byte("TEST"))))
	fmt.Println("ToUpper:   ", string(bytes.ToUpper([]byte("test"))))



	
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.


