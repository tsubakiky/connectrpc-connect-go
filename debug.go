package connect

import (
	"io"
	"os"
	"strings"
	_ "unsafe"

	"github.com/kortschak/utter"
)

// Config is the active configuration of the [utter.Config].
var Config = utter.ConfigState{
	Indent:           " ",
	NumericWidth:     1,
	StringWidth:      1,
	Quoting:          utter.DoubleQuote,
	BytesWidth:       16,
	CommentBytes:     true,
	AddressBytes:     false,
	CommentPointers:  false,
	IgnoreUnexported: false,
	OmitZero:         false,
	ElideType:        false,
	SortKeys:         false,
}

//go:linkname fdump github.com/kortschak/utter.fdump
func fdump(cs *utter.ConfigState, w io.Writer, a any)

// Fdump formats and displays the passed arguments to io.Writer w.
//
// It formats exactly the same as [Dump].
func Fdump(w io.Writer, a any) {
	fdump(&Config, w, a)
}

// Sdump returns a string with the passed arguments formatted exactly the same as [Dump].
func Sdump(a any) string {
	var s strings.Builder

	fdump(&Config, &s, a)
	return s.String()
}

// Dump displays the passed parameters to standard out with newlines, customizable
// indentation, and additional debug information such as complete types and all
// pointer addresses used to indirect to the final value.  It provides the
// following features over the built-in printing facilities provided by the fmt
// package:
//
//   - Pointers are dereferenced and followed
//   - Circular data structures are detected and annotated
//   - Byte arrays and slices are dumped in a way similar to the hexdump -C command,
//     which includes byte values in hex, and ASCII output
//
// The configuration options are controlled by an exported package global,
// [utter.Config]. See ConfigState for options documentation.
//
// See [Fdump] if you would prefer dumping to an arbitrary io.Writer or Sdump to
// get the formatted result as a string.
func Dump(a any) {
	fdump(&Config, os.Stdout, a)
}
