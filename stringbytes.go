// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.
//S1 OMIT
package main

import  ("strings"
	"fmt"
	"bytes"
)
func main() {
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

	//S2 OMIT
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
//S3 OMIT
// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.


