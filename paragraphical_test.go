package paragraphical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SimpleWrap(t *testing.T) {
	text := `This is a fairly long string that needs to be broken in to smaller lines in the terminal.`

	expected20 := `This is a fairly
long string that
needs to be broken
in to smaller lines
in the terminal.`
	assert.Equal(t, expected20, Format(20, text))

	expected30 := `This is a fairly long string
that needs to be broken in to
smaller lines in the terminal.`
	assert.Equal(t, expected30, Format(30, text))
}

func Test_LongWords(t *testing.T) {
	text := `This is another really long sentence with an additionally1234567890123 long word.`

	expected := `This is another
really long sentence
with an
additionally1234567890123
long word.`

	assert.Equal(t, expected, Format(20, text))
}

func Test_Multiline(t *testing.T) {
	text := `This is some more super long text but
this time it's broken over
several lines.

Here is some more text to be wrapped.`

	expected := `This is some more
super long text but
this time it's
broken over
several lines.

Here is some more
text to be wrapped.`

	assert.Equal(t, expected, Format(20, text))
}

func Test_Indentation(t *testing.T) {
	text := `The If-Match HTTP request header makes the request conditional. For GET and HEAD methods, the server will send back the requested resource only if it matches one of the listed ETags. For PUT and other non-safe methods, it will only upload the resource in this case.

The comparison with the stored ETag uses the strong comparison algorithm, meaning two files are considered identical byte to byte only. If a listed ETag has the W/ prefix indicating a weak entity tag, it will never match under this comparison algorithm.

There are two common use cases:

    For GET and HEAD methods, used in combination with a Range header, it can guarantee that the new ranges requested comes from the same resource than the previous one. If it doesn't match, then a 416 (Range Not Satisfiable) response is returned.
    For other methods, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Matchxxxxxxxxxxxxxxxxxxxx`

	expected := `The If-Match HTTP request header makes the request conditional. For GET and HEAD
methods, the server will send back the requested resource only if it matches one
of the listed ETags. For PUT and other non-safe methods, it will only upload the
resource in this case.

The comparison with the stored ETag uses the strong comparison algorithm,
meaning two files are considered identical byte to byte only. If a listed ETag
has the W/ prefix indicating a weak entity tag, it will never match under this
comparison algorithm.

There are two common use cases:

    For GET and HEAD methods, used in combination with a Range header, it can
     guarantee that the new ranges requested comes from the same resource than
     the previous one. If it doesn't match, then a 416 (Range Not Satisfiable)
     response is returned.
    For other methods, and in particular for PUT, If-Match can be used to
     prevent the lost update problem. It can check if the modification of a
     resource that the user wants to upload will not override another change
     that has been done since the original resource was fetched. If the request
     cannot be fulfilled, the 412 (Precondition Failed) response is returned.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Matchxxxxxxxxxxxxxxxxxxxx`

	assert.Equal(t, expected, Format(80, text))
}

func Test_getIndent(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{name: "No indent", parts: []string{"some", "words"}, want: ""},
		{name: "Plain indent", parts: []string{"", "", "", "some", "words"}, want: "   "},
		{name: "Some gaps after indent", parts: []string{"", "", "some", "", "", "words"}, want: "  "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndent(tt.parts); got != tt.want {
				t.Errorf("getIndent(%v) = %v, want %v", tt.parts, got, tt.want)
			}
		})
	}
}
